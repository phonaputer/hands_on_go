package logic

import (
	"errors"
	"fmt"
	"hands_on_go/internal/uaerr"
)

type userServiceImpl struct {
	userRepository UserRepository
}

func NewUserService(
	userRepository UserRepository,
) UserService {
	return &userServiceImpl{
		userRepository: userRepository,
	}
}

func (u *userServiceImpl) GetByID(id int) (*User, error) {
	user, err := u.userRepository.GetUser(id)
	if err != nil {
		return nil, fmt.Errorf("get user: %w", err)
	}

	return user, nil
}

func (u *userServiceImpl) CreateUser(user *User) (int, error) {
	return u.userRepository.CreateUser(user)
}

func (u *userServiceImpl) DeleteByID(id int) error {

	// 1. Check user exists
	exists, err := u.userRepository.CheckUserExists(id)

	// 1.1. check for unexpected error
	if err != nil {
		return fmt.Errorf("check user exists: %w", err)
	}

	// 1.2. return "not found" if not exists
	if !exists {
		err := errors.New("user not found")
		err = uaerr.SetUserMsg(err, "user not found")
		err = uaerr.SetType(err, uaerr.TypeNotFound)

		return err
	}

	// 2. Delete user
	err = u.userRepository.DeleteUser(id)

	// 2.1. check for unexpected error
	if err != nil {
		return fmt.Errorf("delete user: %w", err)
	}

	// 3. return success
	return nil
}
