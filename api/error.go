package api

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/getkin/kin-openapi/openapi3filter"
	"github.com/go-chi/render"
	"github.com/rs/zerolog/hlog"
	"github.com/rs/zerolog/log"
	"github.com/tjdavis3/problems"
)

type ErrResponse problems.Problem

const problemTag = "tag:ringsq.io,2022-05-01:errResponse"

func (e *ErrResponse) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, e.Status)
	return nil
}

func problemFromError(status int, err error) *problems.Problem {
	problem := problems.FromErrorWithStatus(status, err)
	problem.Set("type", problemTag)
	return problem
}

func buildProblemFromError(status int, err error, r *http.Request) *problems.Problem {
	problem := problemFromError(status, err)
	reqID, _ := hlog.IDFromRequest(r)
	problem.Set("instance", r.URL.Path)
	problem.Set("request-id", reqID.String())
	return problem
}

func ErrInvalidRequest(r *http.Request, err error) render.Renderer {
	hlog.FromRequest(r).Error().Err(err).Send()
	problem := buildProblemFromError(http.StatusBadRequest, err, r)
	return problem
}

func ErrNotFound(r *http.Request, err error) render.Renderer {
	hlog.FromRequest(r).Error().Err(err).Send()
	return buildProblemFromError(http.StatusNotFound, err, r)
}

func ErrUnauthorized(r *http.Request, err error) render.Renderer {
	hlog.FromRequest(r).Error().Err(err).Send()
	return buildProblemFromError(http.StatusUnauthorized, err, r)
}

func ErrServerError(r *http.Request, err error) render.Renderer {
	hlog.FromRequest(r).Error().Err(err).Send()
	return buildProblemFromError(http.StatusInternalServerError, err, r)
}

func ErrorEncoder(ctx context.Context, err error, w http.ResponseWriter) {
	log.Ctx(ctx).Error().Err(err).Send()
	reqID, _ := hlog.IDFromCtx(ctx)
	w.Header().Set("Content-Type", "application/json")
	if headerer, ok := err.(openapi3filter.Headerer); ok {
		for k, values := range headerer.Headers() {
			for _, v := range values {
				w.Header().Add(k, v)
			}
		}
	}

	code := http.StatusInternalServerError
	if sc, ok := err.(openapi3filter.StatusCoder); ok {
		code = sc.StatusCode()
	}
	if code == 401 {
		w.Header().Add("WWW-Authenticate", "Bearer")
	}
	problem := problemFromError(code, err)
	problem.Set("request-id", reqID)

	w.WriteHeader(code)
	body, _ := json.Marshal(problem)

	w.Write(body)
}

type JWTErr struct {
	Status int
	Msg    string
}

func (jwte *JWTErr) StatusCode() int {
	return jwte.Status
}
func (jwte *JWTErr) Error() string {
	return jwte.Msg
}

// JWTErrorHandler formatts JWT validation errors with the builtin error
// response.  This is passed to the JWTMiddleware
func JWTErrorHandler(w http.ResponseWriter, r *http.Request, err string) {
	jerr := &JWTErr{Status: http.StatusUnauthorized, Msg: err}
	ctx := r.Context()
	ErrorEncoder(ctx, jerr, w)
	return
}

// BadRequestErrorHandler - used by the server handler for bad requests due to invalid parameters
func BadRequestErrorHandler(w http.ResponseWriter, r *http.Request, err error) {
	problem := buildProblemFromError(http.StatusBadRequest, err, r)
	problem.Render(w, r)
}
