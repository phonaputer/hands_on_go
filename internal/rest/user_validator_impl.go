package rest

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/phonaputer/hands_on_go/internal/blerr"
	"io"
	"net/http"
	"strconv"
)

type UserValidatorImpl struct{}

func (u *UserValidatorImpl) ValidateCreateUser(r *http.Request) (*createUserRequest, error) {
	var req createUserRequest
	var err error

	err = validJSONBody(err, r.Body, &req)
	err = required(err, req.FirstName, "firstName")
	err = required(err, req.LastName, "lastName")
	err = required(err, req.Age, "age")
	err = required(err, req.PhoneNumber, "phoneNumber")
	err = required(err, req.IsPhoneNumberVerified, "isPhoneNumberVerified")
	err = jsonNumberIsInt(err, req.Age, "age")
	err = stringLength(err, 1, 100, req.FirstName, "firstName")
	err = stringLength(err, 1, 100, req.LastName, "lastName")
	err = stringLength(err, 1, 25, req.PhoneNumber, "phoneNumber")
	if err != nil {
		return nil, fmt.Errorf("%s: %w", err, blerr.ErrInvalidInput)
	}

	return &req, nil
}

func (u *UserValidatorImpl) ValidateGetUserByID(r *http.Request) (int, error) {
	id, err := intPathVariable(nil, r, "id")
	if err != nil {
		return 0, fmt.Errorf("%e: %w", err, blerr.ErrInvalidInput)
	}

	return id, nil
}

func (u *UserValidatorImpl) ValidateDeleteUserByID(r *http.Request) (int, error) {
	id, err := intPathVariable(nil, r, "id")
	if err != nil {
		return 0, fmt.Errorf("%e: %w", err, blerr.ErrInvalidInput)
	}

	return id, nil
}

func intPathVariable(err error, r *http.Request, pathVar string) (int, error) {
	varStr, ok := mux.Vars(r)[pathVar]

	if !ok {
		return 0, fmt.Errorf("path variable '%s' is required", pathVar)
	}

	result, err := strconv.Atoi(varStr)
	if err != nil {
		return 0, fmt.Errorf("path variable '%s' has invalid format", pathVar)
	}

	return result, nil
}

func validJSONBody(err error, r io.Reader, result interface{}) error {
	if err != nil {
		return err // "short circuit" if the error is already not nil
	}

	err = json.NewDecoder(r).Decode(result)
	if err != nil {
		return errors.New("invalid json body")
	}

	return nil
}

func required[T any](err error, value *T, valueName string) error {
	if err != nil {
		return err // "short circuit" if the error is already not nil
	}

	if value == nil {
		return fmt.Errorf("%s is required", valueName)
	}

	return nil
}

func stringLength(err error, min int, max int, value *string, valueName string) error {
	if err != nil || value == nil {
		return err
	}

	length := len(*value)
	if length < min || max < length {
		return fmt.Errorf("%s has invalid length", valueName)
	}

	return nil
}

func jsonNumberIsInt(err error, value *json.Number, valueName string) error {
	if err != nil || value == nil {
		return err
	}

	_, err = value.Int64()
	if err != nil {
		return fmt.Errorf("%s is not an integer", valueName)
	}

	return nil
}
