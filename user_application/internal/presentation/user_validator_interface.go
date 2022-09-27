package presentation

import (
	"errors"
	"net/http"
)

var errInvalidInput = errors.New("invalid input")

type userValidator interface {
	// ValidateGetUserByID returns errInvalidInput if the request data is not valid.
	ValidateGetUserByID(r *http.Request) (int, error)
}
