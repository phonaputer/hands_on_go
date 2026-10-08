package logic

import "context"

type UserRepository interface {
	GetByID(ctx context.Context, id int) (*User, error)
}
