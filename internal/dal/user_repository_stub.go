package dal

import (
	"errors"
	"github.com/phonaputer/hands_on_go/internal/blerr"
	"github.com/phonaputer/hands_on_go/internal/model"
)

type UserRepositoryStub struct{}

func (u *UserRepositoryStub) CheckExistsByID(userId int) (bool, error) {
	if userId == 501 {
		return false, errors.New("whoa!")
	}

	return userId != 404, nil
}

func (u *UserRepositoryStub) DeleteByID(userId int) error {
	if userId == 500 {
		return errors.New("whoa!")
	}

	return nil
}

func (u *UserRepositoryStub) GetByID(userId int) (*model.User, error) {
	if userId == 404 {
		return nil, blerr.SetKind(errors.New("notfound"), blerr.KindNotFound)
	}
	if userId == 502 {
		return nil, errors.New("whoa!")
	}

	return &model.User{
		ID:              123,
		FirstName:       "first",
		LastName:        "last",
		Age:             456,
		PhoneNumber:     "555-5555",
		IsPhoneVerified: true,
	}, nil
}

func (u *UserRepositoryStub) Create(user *model.User) (int, error) {
	if user.FirstName == "500" {
		return 0, errors.New("whoa!")
	}

	return 123, nil
}
