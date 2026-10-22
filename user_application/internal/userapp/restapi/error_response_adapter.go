package restapi

import (
	"fmt"
	"net/http"
)

type errHTTPHandlerFunc func(w http.ResponseWriter, r *http.Request) error

type errorResponseAdapter struct{}

func newErrorResponseAdapter() *errorResponseAdapter {
	return &errorResponseAdapter{}
}

func (e *errorResponseAdapter) WithErrResp(next errHTTPHandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := next(w, r)
		if err == nil {
			return
		}

		fmt.Println("error handling request:" + err.Error())

		w.WriteHeader(http.StatusInternalServerError)
	}
}
