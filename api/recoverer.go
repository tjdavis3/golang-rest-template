package api

import (
	"errors"
	"net/http"

	sentry "github.com/getsentry/sentry-go"
	"github.com/go-chi/render"
	"github.com/rs/zerolog/hlog"
)

// EventEnhancer is an http middleware to add the request ID to the sentry tags
func EventEnhancer(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		if hub := sentry.GetHubFromContext(ctx); hub != nil {
			var req_id = "unknown"
			if xid, ok := hlog.IDFromCtx(ctx); ok {
				req_id = xid.String()
			}
			hub.Scope().SetTag("RequestID", req_id)
		}
		next.ServeHTTP(w, r.WithContext(ctx))
	}
	return http.HandlerFunc(fn)

}

// Recoverer is a middleware that recovers from panics, logs the panic (and a
// backtrace), and returns a HTTP 500 (Internal Server Error) status if
// possible. Recoverer prints a request ID if one is provided.
//
// Alternatively, look at https://github.com/pressly/lg middleware pkgs.
func Recoverer(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if rvr := recover(); rvr != nil && rvr != http.ErrAbortHandler {
				var rvrerr error
				if err, ok := rvr.(error); ok {
					rvrerr = err
				} else if err, ok := rvr.(string); ok {
					rvrerr = errors.New(err)
				} else {
					rvrerr = errors.New("unknown error")
				}

				render.Render(w, r, ErrServerError(r, rvrerr))

			}
		}()

		next.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}
