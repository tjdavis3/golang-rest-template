package api

import (
	"context"
	"net/http"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/getkin/kin-openapi/openapi3filter"
)

type AuthenticationFunc func(context.Context, *openapi3filter.AuthenticationInput) error

func NoopAuthenticationFunc(context.Context, *openapi3filter.AuthenticationInput) error { return nil }

var _ AuthenticationFunc = NoopAuthenticationFunc

type ValidationHandler struct {
	Handler            http.Handler
	AuthenticationFunc AuthenticationFunc
	Swagger            *openapi3.Swagger
	ErrorEncoder       openapi3filter.ErrorEncoder
	router             *openapi3filter.Router
}

func NewValidationHandler(handler http.Handler, swagger *openapi3.Swagger) (*ValidationHandler, error) {
	vh := &ValidationHandler{Handler: handler, Swagger: swagger}
	err := vh.Load()
	return vh, err
}

func (h *ValidationHandler) Load() error {
	h.router = openapi3filter.NewRouter()

	if err := h.router.AddSwagger(h.Swagger); err != nil {
		return err
	}

	// set defaults
	if h.Handler == nil {
		h.Handler = http.DefaultServeMux
	}
	if h.AuthenticationFunc == nil {
		h.AuthenticationFunc = NoopAuthenticationFunc
	}
	if h.ErrorEncoder == nil {
		encoder := &openapi3filter.ValidationErrorEncoder{Encoder: (openapi3filter.ErrorEncoder)(ErrorEncoder)}
		h.ErrorEncoder = encoder.Encode
	}

	return nil
}

func (h *ValidationHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if handled := h.before(w, r); handled {
		return
	}
	// TODO: validateResponse
	h.Handler.ServeHTTP(w, r)
}

// Middleware implements gorilla/mux MiddlewareFunc
func (h *ValidationHandler) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if handled := h.before(w, r); handled {
			return
		}
		// TODO: validateResponse
		next.ServeHTTP(w, r)
	})
}

func (h *ValidationHandler) before(w http.ResponseWriter, r *http.Request) (handled bool) {
	if err := h.validateRequest(r); err != nil {
		h.ErrorEncoder(r.Context(), err, w)
		return true
	}
	return false
}

func (h *ValidationHandler) validateRequest(r *http.Request) error {
	// Find route
	route, pathParams, err := h.router.FindRoute(r.Method, r.URL)
	if err != nil {
		return err
	}

	options := &openapi3filter.Options{
		AuthenticationFunc: h.AuthenticationFunc,
	}

	// Validate request
	requestValidationInput := &openapi3filter.RequestValidationInput{
		Request:    r,
		PathParams: pathParams,
		Route:      route,
		Options:    options,
	}
	if err = openapi3filter.ValidateRequest(r.Context(), requestValidationInput); err != nil {
		return err
	}

	return nil
}
