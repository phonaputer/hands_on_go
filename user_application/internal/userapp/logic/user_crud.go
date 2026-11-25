package logic

import (
	"context"
	"fmt"
)

type UserCRUD interface {
	GetByID(ctx context.Context, userID int) (*User, error)
	Create(ctx context.Context, user *NewUser) (int, error)
}

type NewUser struct {
	FirstName  string
	MiddleName *string
	LastName   string
	Age        int
	Email      string
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

func (c *UserCRUDImpl) Create(ctx context.Context, user *NewUser) (int, error) {
	exists, err := c.repo.CheckExistsByEmail(ctx, user.Email)
	if err != nil {
		return 0, fmt.Errorf("check exists by email: %w", err)
	}

	if exists {
		return 0, ErrEmailAlreadyRegistered
	}

	id, err := c.repo.Create(ctx, user)
	if err != nil {
		return 0, fmt.Errorf("create: %w", err)
	}

	return id, nil
}
