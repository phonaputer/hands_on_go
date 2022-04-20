package rest

import (
	"github.com/phonaputer/hands_on_go/internal/blerr"
	"net/http"
)

type UserValidatorStubImpl struct{}

func (u *UserValidatorStubImpl) ValidateGetUserByID(r *http.Request) (int, error) {

	if r.URL.Query().Has("invalid") {
		return 0, blerr.ErrInvalidInput
	}
	if r.URL.Query().Has("notfound") {
		return 404, nil
	}

	return 1, nil
}

func (u *UserValidatorStubImpl) ValidateDeleteUserByID(r *http.Request) (int, error) {

	if r.URL.Query().Has("invalid") {
		return 0, blerr.ErrInvalidInput
	}
	if r.URL.Query().Has("notfound") {
		return 404, nil
	}

	return 1, nil
}
