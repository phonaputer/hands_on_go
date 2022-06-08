package logic

import "github.com/phonaputer/hands_on_go/internal/model"

type UserRepository interface {
	CheckExistsByID(userId int) (bool, error)
	DeleteByID(userId int) error
	GetByID(userId int) (*model.User, error)
	Create(user *model.User) (int, error)
}
