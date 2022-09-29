package presentation

import (
	"errors"
	"net/http"
)

var errInvalidInput = errors.New("invalid input")

// userValidator validates request data for the /users endpoints
//
// All functions return errInvalidInput if the request data is not valid.
type userValidator interface {
	ValidateGetUserByID(r *http.Request) (int, error)
	ValidateCreateUser(r *http.Request) (*createUserRequest, error)
	ValidateDeleteUserByID(r *http.Request) (int, error)
}
