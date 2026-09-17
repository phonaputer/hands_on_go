package logic

import (
	"context"
	"time"
)

type UserCRUD interface {
	GetByID(ctx context.Context, userID int) (*User, error)
}

type DummyUserCRUD struct{}

func NewDummyUserCRUD() DummyUserCRUD {
	return DummyUserCRUD{}
}

func (c DummyUserCRUD) GetByID(ctx context.Context, userID int) (*User, error) {
	return &User{
		ID:         userID,
		FirstName:  "test",
		MiddleName: new("middle"),
		LastName:   "last",
		Age:        123,
		Email:      "123@example.com",
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now().Add(time.Minute),
	}, nil
}
