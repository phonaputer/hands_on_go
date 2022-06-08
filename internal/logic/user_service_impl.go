package logic

import (
	"errors"
	"fmt"
	"github.com/phonaputer/hands_on_go/internal/blerr"
	"github.com/phonaputer/hands_on_go/internal/model"
)

type UserServiceImpl struct {
	userRepository UserRepository
}

func NewUserServiceImpl(userRepository UserRepository) *UserServiceImpl {
	return &UserServiceImpl{
		userRepository: userRepository,
	}
}

func (u *UserServiceImpl) GetByID(userID int) (*model.User, error) {
	result, err := u.userRepository.GetByID(userID)
	if blerr.GetKind(err) == blerr.KindNotFound {
		return nil, blerr.SetUserMsg(err, "user not found")
	}
	if err != nil {
		return nil, fmt.Errorf("get from user repository: %w", err)
	}

	return result, nil
}

func (u *UserServiceImpl) DeleteByID(userID int) error {
	exists, err := u.userRepository.CheckExistsByID(userID)
	if err != nil {
		return fmt.Errorf("check user exists: %w", err)
	}
	if !exists {
		err := blerr.SetKind(errors.New("user not found"), blerr.KindNotFound)
		return blerr.SetUserMsg(err, "user not found")
	}

	err = u.userRepository.DeleteByID(userID)
	if err != nil {
		return fmt.Errorf("delete user: %w", err)
	}

	return nil
}

func (u *UserServiceImpl) Create(user *model.User) (int, error) {
	id, err := u.userRepository.Create(user)
	if err != nil {
		return 0, fmt.Errorf("creating user: %w", err)
	}
	return id, nil
}
