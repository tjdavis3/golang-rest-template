//go:generate oapi-codegen --package=api --generate types,chi-server,spec -o api.gen.go ../spec/openapi.yaml
//go:generate go run ../cmd/updImpl.go

package api

import (
	"context"
	"errors"
	"io"
	"net/http"
	"time"

	"../models"
	sentryhttp "github.com/getsentry/sentry-go/http"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/hlog"
	"github.com/rs/zerolog/log"
)

type server struct {
	db         *interface{}
	httpClient *http.Client
	config     *Cfg
	root       http.Handler
}

func NewServer(ctx context.Context, cfg *Cfg) (*server, error) {
	var err error

	s := &server{
		config: cfg,
	}

	zerolog.SetGlobalLevel(zerolog.DebugLevel)
	log.Logger = log.Logger.With().Timestamp().Caller().Logger()

	// db
	s.db, err = models.InitializeDB(s.config)
	if err != nil {
		return nil, err
	}

	// configure http client for global usage
	s.httpClient = &http.Client{
		Timeout: time.Second * 10,
	}

	// routers, middlewares
	r := chi.NewRouter()
	sentryHandler := sentryhttp.New(sentryhttp.Options{
		Repanic:         true,
		WaitForDelivery: true,
		// Timeout for the event delivery requests.
		Timeout: 3})

	// TODO: Add jwt-go-middleware to validate JWT.  For the ValidationKeyGetter see https://auth0.com/docs/quickstart/backend/golang/01-authorization
	// jwtMiddleware := jwtmiddleware.New(jwtmiddleware.Options{
	// 	ErrorHandler:        JWTErrorHandler,
	// 	CredentialsOptional: false,
	// 	Debug:               true,
	// 	ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
	// 		secret := "MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA7/mAP0uVZQcYC3JtDCjelfZwqNp3kbOsBG2d2ILcDIEUEMs2VnTgTaMHky2/3dLF/wVYpD3ObNquIJslwdxrxxyXBoNKEkzdI34UgjB+ZcX7S++THLyg7bAkEMAn9jGK3wnPHpgK5Karxnu5dCBU6QPocekWeu5ibQr8gnxiaR4WdsYZhaRwRD0VvH1kSOtx2ceYnmtRACJv3MtPraJxUmVsa7Yzu8GRCd+EqeKRMkX/p8hNCdws04t9dO3AemVGI2gGAwJ3d16yPNd0hFWFOF58CVTD6fyDqPqE74DhSzBrJmggEkaxehLUpUofvP9WPrQz8YsDyMGjqByun+8VHQIDAQAB"
	// 		return []byte(secret), nil
	// 	},
	// 	SigningMethod: jwt.SigningMethodRS256,
	// })

	r.Handle("/metrics", promhttp.Handler())

	r.Group(func(r chi.Router) {
		// If you are service is behind load balancer like nginx, you might want to
		// use X-Request-ID instead of injecting request id. You can do some thing
		// like this,
		// r.Use(hlog.CustomHeaderHandler("reqId", "X-Request-Id"))
		r.Use(middleware.Recoverer)
		r.Use(hlog.NewHandler(log.Logger))
		r.Use(hlog.AccessHandler(func(r *http.Request, status, size int, duration time.Duration) {
			hlog.FromRequest(r).Info().
				Str("method", r.Method).
				Str("url", r.URL.String()).
				Int("status", status).
				Int("size", size).
				Dur("duration", duration).
				Msg("")
		}))
		r.Use(hlog.RequestIDHandler("req_id", "Request-Id"))
		r.Use(hlog.RemoteAddrHandler("ip"))
		r.Use(hlog.UserAgentHandler("user_agent"))
		r.Use(hlog.RefererHandler("referer"))
		r.Use(mwMetrics)
		r.Use(Recoverer)

		r.Use(s.JWTAuthentication)
		r.Use(sentryHandler.Handle)
		r.Use(EventEnhancer)
		handler := HandlerFromMux(s, r)
		r.Handle("/", handler)
	})

	r.HandleFunc("/", notFoundHandler)

	// health check
	r.HandleFunc("/ping", ping)
	r.HandleFunc("/openapi.json", spec)

	s.root = r

	return s, nil
}

func spec(w http.ResponseWriter, r *http.Request) {
	swagger, err := GetSwagger()
	if err != nil {
		render.Render(w, r, ErrServerError(r, err))
		return
	}
	render.JSON(w, r, swagger)
}

func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	render.Render(w, r, ErrNotFound(r, errors.New("Not Found")))
}

// ping is handler responding to health-check request
func ping(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "pong")
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.root.ServeHTTP(w, r)
}
