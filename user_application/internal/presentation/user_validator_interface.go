package presentation

import (
	"net/http"
)

// userValidator validates request data for the /users endpoints
//
// All functions return uaerr.TypeInvalidInput if the request data is not valid.
type userValidator interface {
	ValidateGetUserByID(r *http.Request) (int, error)
	ValidateCreateUser(r *http.Request) (*createUserRequest, error)
	ValidateDeleteUserByID(r *http.Request) (int, error)
}
