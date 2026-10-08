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
		FirstName:  "asfsf",
		MiddleName: nil,
		LastName:   "asdgs",
		Age:        122,
		Email:      "dasgsdaga",
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}, nil
}
