package rest

import "net/http"

type userValidator interface {
	// ValidateGetUserByID returns a User’s integer ID or an error
	ValidateGetUserByID(r *http.Request) (int, error)

	// ValidateDeleteUserByID returns a User’s integer ID or an error
	ValidateDeleteUserByID(r *http.Request) (int, error)

	ValidateCreateUser(r *http.Request) (*createUserRequest, error)
}
