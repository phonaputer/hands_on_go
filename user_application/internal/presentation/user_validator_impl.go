package presentation

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
	"golang.org/x/exp/constraints"
	"io"
	"net/http"
	"strconv"
)

type UserValidatorImpl struct{}

func (u *UserValidatorImpl) ValidateDeleteUserByID(r *http.Request) (int, error) {
	return u.ValidateGetUserByID(r)
}

func (u *UserValidatorImpl) ValidateGetUserByID(r *http.Request) (int, error) {

	// 1. does query string contain "id"
	if !r.URL.Query().Has("id") {
		return 0, errInvalidInput // TODO add a better message
	}

	// 2. is query string "id" an integer
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		logrus.WithError(err).Trace("error parsing user id from query string")

		return 0, errInvalidInput // TODO add a better message
	}

	return id, nil
}

func (u *UserValidatorImpl) ValidateCreateUser(r *http.Request) (*createUserRequest, error) {

	var err error
	var result createUserRequest

	// 1. parse JSON
	err = validJSONBody(err, r.Body, &result)

	// 2. validate required fields
	err = required(err, result.FirstName, "firstName")
	err = required(err, result.LastName, "lastName")
	err = required(err, result.Age, "age")
	err = required(err, result.PhoneNumber, "phoneNumber")
	err = required(err, result.IsPhoneVerified, "isPhoneVerified")

	// 3. validate string field lengths
	err = validateStringLength(err, result.PhoneNumber, 1, 25, "phoneNumber")
	err = validateStringLength(err, result.FirstName, 1, 100, "firstName")
	err = validateStringLength(err, result.LastName, 1, 100, "lastName")

	// 4. validate int size of age field
	err = validateSize(err, result.Age, 1, 200, "age")

	// 5.A. return error result if there was a validation problem
	if err != nil {
		return nil, fmt.Errorf("%e : %w", err, errInvalidInput)
	}

	// 5.B. return success result
	return &result, nil
}

func validJSONBody(err error, r io.Reader, result interface{}) error {
	if err != nil {
		return err // "short circuit" if the error is already not nil
	}

	err = json.NewDecoder(r).Decode(result)
	if err != nil {
		logrus.WithError(err).Debug("error parsing JSON for create user")

		return errors.New("invalid json body")
	}

	return nil
}

func validateStringLength(err error, value *string, min int, max int, fieldName string) error {
	if err != nil || value == nil {
		return err
	}

	valueLen := len(*value)
	if valueLen < min || max < valueLen {
		return fmt.Errorf("%s has invalid length", fieldName)
	}

	return nil
}

func validateSize[T constraints.Ordered](err error, value *T, min T, max T, fieldName string) error {
	if err != nil || value == nil {
		return err
	}

	if *value < min || max < *value {
		return fmt.Errorf("%s has invalid length", fieldName)
	}

	return nil
}

func required[T any](err error, value *T, fieldName string) error {
	if err != nil {
		return err
	}

	if value == nil {
		return fmt.Errorf("%s is required", fieldName)
	}

	return nil
}
