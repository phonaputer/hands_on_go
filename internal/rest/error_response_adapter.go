package rest

import (
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
	switch blerr.GetKind(err) {
	case blerr.KindInvalidInput:
		return 400, userMsgOrDefault(err, "invalid input")
	case blerr.KindNotFound:
		return 404, userMsgOrDefault(err, "not found")
	default:
		return 500, "internal server error"
	}
}

func userMsgOrDefault(err error, defaultMsg string) string {
	if userMsg, ok := blerr.GetUserMsg(err); ok {
		return userMsg
	}

	return defaultMsg
}
