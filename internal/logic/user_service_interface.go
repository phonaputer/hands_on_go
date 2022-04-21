package logic

import "github.com/phonaputer/hands_on_go/internal/model"

type UserService interface {
	// GetByID returns blerr.ErrUserNotFound if the user is missing
	GetByID(userID int) (*model.User, error)

	// DeleteByID returns blerr.ErrUserNotFound if the user is missing
	DeleteByID(userID int) error

	// Create returns the new users's unique integer ID if the user is successfully created.
	Create(user *model.User) (int, error)
}
