package logic

type UserRepository interface {
	CreateUser(user *User) (int, error)

	GetUser(id int) (*User, error)

	// CheckUserExists returns false if the user does not exist and true otherwise.
	CheckUserExists(id int) (bool, error)

	// DeleteUser will return an error if the user to be deleted does not exist.
	// Otherwise the user will be deleted.
	DeleteUser(id int) error
}
