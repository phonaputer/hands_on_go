package logic

import "errors"

var ErrNotFound = errors.New("not found")

type UserService interface {

	// GetByID returns ErrNotFound if a user with this ID is not found.
	GetByID(id int) (*User, error)

	// CreateUser returns the ID of the newly created user.
	CreateUser(user *User) (int, error)
}
