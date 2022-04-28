package logic

import (
	"errors"
	"github.com/phonaputer/hands_on_go/internal/blerr"
	"github.com/phonaputer/hands_on_go/internal/model"
)

type UserServiceStubImpl struct{}

func (u *UserServiceStubImpl) GetByID(userID int) (*model.User, error) {
	if userID == 404 {
		err := blerr.SetKind(errors.New("test"), blerr.KindNotFound)
		return nil, blerr.SetUserMsg(err, "user not found")
	}

	return &model.User{FirstName: "Risa", LastName: "Rakuten", Age: 50,
			PhoneNumber: "1-800-555-5555", IsPhoneVerified: true},
		nil
}

func (u *UserServiceStubImpl) DeleteByID(userID int) error {
	if userID == 404 {
		err := blerr.SetKind(errors.New("test"), blerr.KindNotFound)
		return blerr.SetUserMsg(err, "user not found")
	}

	return nil
}

func (u *UserServiceStubImpl) Create(user *model.User) (int, error) {
	return 1, nil
}
