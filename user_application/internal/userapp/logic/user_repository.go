package logic

import "context"

type UserRepository interface {
	GetByID(ctx context.Context, id int) (*User, error)
	CheckExistsByEmail(ctx context.Context, email string) (bool, error)
	Create(ctx context.Context, user *NewUser) (int, error)
}
