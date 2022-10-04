package presentation

import (
	"github.com/sirupsen/logrus"
	"hands_on_go/internal/uaerr"
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
	errType := uaerr.GetType(err)

	switch errType {
	case uaerr.TypeInvalidInput:
		return 400, getUserMsgOrDefault(err, "invalid request")

	case uaerr.TypeNotFound:
		return 404, getUserMsgOrDefault(err, "not found")
	default:
		return 500, "internal server error"
	}
}

func getUserMsgOrDefault(err error, defaultMsg string) string {
	msg, hasMsg := uaerr.GetUserMsg(err)

	if hasMsg {
		return msg
	}

	return defaultMsg
}
