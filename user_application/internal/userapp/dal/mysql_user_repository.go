package dal

import (
	"context"
	"time"
	"userapp/internal/userapp/logic"
)

type MySQLUserRepository struct {
}

func newMySQLUserRepository() *MySQLUserRepository {
	return &MySQLUserRepository{}
}

func (r *MySQLUserRepository) GetByID(ctx context.Context, id int) (*logic.User, error) {
	return &logic.User{
		ID:         id,
		FirstName:  "test",
		MiddleName: nil,
		LastName:   "test",
		Age:        123,
		Email:      "test",
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}, nil
}
