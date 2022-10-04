package logic

type UserService interface {

	// GetByID returns uaerr.TypeNotFound if a user with this ID is not found.
	GetByID(id int) (*User, error)

	// CreateUser returns the ID of the newly created user.
	CreateUser(user *User) (int, error)

	// DeleteByID returns ErrNotFound if a user with this ID is not found.
	DeleteByID(id int) error
}
