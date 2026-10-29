package restapi

import (
	"errors"
	"fmt"
	"net/http"
	"userapp/internal/userapp/logic"
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

		e.writeRespnse(err, w)
	}
}

func (e *errorResponseAdapter) writeRespnse(
	err error,
	w http.ResponseWriter,
) {
	switch {
	case errors.Is(err, logic.ErrUserNoFound):
		w.WriteHeader(http.StatusNotFound)
	case errors.Is(err, logic.ErrInvalidInput):
		w.WriteHeader(http.StatusBadRequest)
	default:
		w.WriteHeader(http.StatusInternalServerError)
	}
}
