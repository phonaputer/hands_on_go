package rest

import (
	"errors"
	"github.com/phonaputer/hands_on_go/internal/blerr"
	"github.com/sirupsen/logrus"
	"net/http"
)

type ErrHTTPHandler func(w http.ResponseWriter, r *http.Request) error

func ErrResponseAdapter(next ErrHTTPHandler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		err := next(w, r)
		if err != nil {
			writeErrorResponse(w, r, err)
		}
	})
}

func writeErrorResponse(w http.ResponseWriter, r *http.Request, err error) {
	logrus.WithError(err).Warnf("error handling request: %s %s", r.Method, r.RequestURI)

	status, message := getStatusAndMsg(err)

	w.WriteHeader(status)
	w.Write([]byte(message))
}

func getStatusAndMsg(err error) (int, string) {
	if errors.Is(err, blerr.ErrInvalidInput) {
		return 400, "request not valid"
	}
	if errors.Is(err, blerr.ErrUserNotFound) {
		return 404, "user not found"
	}

	return 500, "internal server error"
}
