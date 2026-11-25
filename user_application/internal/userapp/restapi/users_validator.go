package restapi

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"userapp/internal/userapp/logic"
)

type usersValidator interface {
	Create(r *http.Request) (*logic.NewUser, error)
}

type jsonUsersValidator struct{}

func newJSONUsersValidator() jsonUsersValidator {
	return jsonUsersValidator{}
}

func (v jsonUsersValidator) Create(r *http.Request) (*logic.NewUser, error) {
	var body createUserRequest

	err := json.NewDecoder(r.Body).Decode(&body)

	err = notNil(err, body.FirstName, "first_name is required")
	err = notNil(err, body.LastName, "last_name is required")
	err = notNil(err, body.Age, "age is required")
	err = notNil(err, body.Email, "email is required")

	err = stringLength(err, body.FirstName, 1, 100, "first_name has invalid length")
	err = stringLength(err, body.MiddleName, 1, 100, "middle_name has invalid length")
	err = stringLength(err, body.LastName, 1, 100, "last_name has invalid length")
	err = stringLength(err, body.Email, 1, 255, "email has invalid length")

	// 3. validate that age < 100

	err = intSize(err, body.Age, 0, 99, "age must be between 0 and 99")

	if err != nil {
		err = logic.SetEndUserMsg(err, err.Error())
		return nil, fmt.Errorf("%w: %w", logic.ErrInvalidInput, err)
	}

	return &logic.NewUser{
		FirstName:  *body.FirstName,
		MiddleName: body.MiddleName,
		LastName:   *body.LastName,
		Age:        *body.Age,
		Email:      *body.Email,
	}, nil
}

func intSize(err error, value *int, min, max int, msg string) error {
	if err != nil {
		return err
	}
	if value == nil {
		return err
	}

	if max < *value || *value < min {
		return errors.New(msg)
	}

	return nil
}

func stringLength(err error, value *string, min, max int, msg string) error {
	if err != nil {
		return err
	}
	if value == nil {
		return err
	}

	length := len([]rune(*value))

	if max < length || length < min {
		return errors.New(msg)
	}

	return nil
}

func notNil[T any](err error, value *T, msg string) error {
	if err != nil {
		return err
	}

	if value == nil {
		return errors.New(msg)
	}

	return nil
}
