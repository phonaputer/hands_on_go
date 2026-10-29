package dal

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"
	"userapp/internal/userapp/logic"

	"github.com/jmoiron/sqlx"
)

type MySQLUserRepository struct {
	db *sqlx.DB
}

func newMySQLUserRepository(
	db *sqlx.DB,
) *MySQLUserRepository {
	return &MySQLUserRepository{
		db: db,
	}
}

type userMySQLRow struct {
	ID         int       `db:"id"`
	FirstName  string    `db:"first_name"`
	MiddleName *string   `db:"middle_name"`
	LastName   string    `db:"last_name"`
	Age        int       `db:"age"`
	Email      string    `db:"email"`
	CreatedAt  time.Time `db:"created_at"`
	UpdatedAt  time.Time `db:"updated_at"`
}

func (r *MySQLUserRepository) GetByID(ctx context.Context, id int) (*logic.User, error) {
	const query = `SELECT * FROM users WHERE id=?`

	var row userMySQLRow
	err := r.db.QueryRowxContext(ctx, query, id).StructScan(&row)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, logic.ErrUserNotFound
	}
	if err != nil {
		return nil, fmt.Errorf("select: %w", err)
	}

	return &logic.User{
		ID:         row.ID,
		FirstName:  row.FirstName,
		MiddleName: row.MiddleName,
		LastName:   row.LastName,
		Age:        row.Age,
		Email:      row.Email,
		CreatedAt:  row.CreatedAt,
		UpdatedAt:  row.UpdatedAt,
	}, nil
}
