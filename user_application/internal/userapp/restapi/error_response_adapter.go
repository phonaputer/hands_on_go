package restapi

import (
	"errors"
	"log/slog"
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

		e.writeResponse(err, w, r)
	}
}

func (e *errorResponseAdapter) writeResponse(
	err error,
	w http.ResponseWriter,
	r *http.Request,
) {
	switch {
	case errors.Is(err, logic.ErrUserNotFound):
		w.WriteHeader(http.StatusNotFound)

	case errors.Is(err, logic.ErrInvalidInput):
		w.WriteHeader(http.StatusBadRequest)

	default:
		w.WriteHeader(http.StatusInternalServerError)
	}

	msg, ok := logic.GetEndUserMsg(err)
	if ok {
		w.Write([]byte(msg + "\n"))
	}

	slog.Error(
		"Error handling request",
		slog.String("path", r.Pattern),
		slog.String("err", err.Error()),
	)
}
