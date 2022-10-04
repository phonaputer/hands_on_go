package presentation

import (
	"errors"
	"github.com/sirupsen/logrus"
	"hands_on_go/internal/logic"
	"net/http"
)

type ErrHTTPHandler func(w http.ResponseWriter, r *http.Request) error

func WithErrResponse(next ErrHTTPHandler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		err := next(w, r)
		if err != nil {
			writeErrorResponse(w, r, err)
		}
	})
}

func writeErrorResponse(w http.ResponseWriter, r *http.Request, err error) {
	logrus.WithError(err).Errorf("an error occurred while serving %s %s", r.Method, r.RequestURI)

	code, msg := getCodeAndMessage(err)

	w.WriteHeader(code)
	w.Write([]byte(msg))
}

func getCodeAndMessage(err error) (int, string) {
	if errors.Is(err, logic.ErrNotFound) {
		return 404, "not found"
	}
	if errors.Is(err, errInvalidInput) {
		return 400, "invalid request"
	}

	return 500, "internal server error"
}
