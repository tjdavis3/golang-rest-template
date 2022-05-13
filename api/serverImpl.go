package api

import (
	"net/http"
)

func (s *server) GetTest(w http.ResponseWriter, r *http.Request, params GetTestParams) {
	err := &ErrResponse{HTTPStatusCode: http.StatusNotImplemented}
	err.Render(w, r)
}
