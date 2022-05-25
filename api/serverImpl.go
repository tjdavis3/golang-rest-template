package api

import (
	"fmt"
	"net/http"
)

var errNotImplemented = fmt.Errorf("This method has not been implemented yet")

func (s *server) GetTest(w http.ResponseWriter, r *http.Request, params GetTestParams) {
	err := buildProblemFromError(http.StatusNotImplemented, errNotImplemented, r)
	err.Render(w, r)
}
