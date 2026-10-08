package logic

import (
	"context"
)

type UserCRUD interface {
	GetByID(ctx context.Context, userID int) (*User, error)
}

type UserCRUDImpl struct {
	repo UserRepository
}

func newUserCRUDImpl(
	repo UserRepository,
) *UserCRUDImpl {
	return &UserCRUDImpl{
		repo: repo,
	}
}

func (c *UserCRUDImpl) GetByID(ctx context.Context, userID int) (*User, error) {
	return c.repo.GetByID(ctx, userID)
}
