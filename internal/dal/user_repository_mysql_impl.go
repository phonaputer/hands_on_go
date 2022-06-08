package dal

import (
	"database/sql"
	"fmt"
	"github.com/phonaputer/hands_on_go/internal/blerr"
	"github.com/phonaputer/hands_on_go/internal/model"
)

type UserRepositoryMySQLImpl struct {
	db *sql.DB
}

func NewUserRepositoryMySQLImpl(db *sql.DB) *UserRepositoryMySQLImpl {
	return &UserRepositoryMySQLImpl{
		db: db,
	}
}

func (u *UserRepositoryMySQLImpl) CheckExistsByID(userId int) (bool, error) {
	const query = `SELECT count(*) FROM users WHERE id=?`

	row := u.db.QueryRow(query, userId)

	var count int
	err := row.Scan(&count)
	if err != nil {
		return false, fmt.Errorf("SELECT count(*): %w", err)
	}

	return count > 0, nil
}

func (u *UserRepositoryMySQLImpl) DeleteByID(userId int) error {
	const query = `DELETE FROM users WHERE id=?`

	_, err := u.db.Exec(query, userId)
	if err != nil {
		return fmt.Errorf("DELETE: %w", err)
	}

	return nil
}

func (u *UserRepositoryMySQLImpl) GetByID(userId int) (*model.User, error) {
	const query = `SELECT id, first_name, last_name, age, phone_number, phone_verification_status
						FROM users WHERE id=?`

	row := u.db.QueryRow(query, userId)

	var result model.User

	err := row.Scan(&result.ID, &result.FirstName, &result.LastName, &result.Age,
		&result.PhoneNumber, &result.IsPhoneVerified)
	if err == sql.ErrNoRows {
		return nil, blerr.SetKind(err, blerr.KindNotFound)
	}
	if err != nil {
		return nil, fmt.Errorf("SELECT count(*): %w", err)
	}

	return &result, nil
}

func (u *UserRepositoryMySQLImpl) Create(user *model.User) (int, error) {
	const query = `INSERT INTO users 
						(first_name, last_name, age, phone_number, phone_verification_status)
						VALUES (?, ?, ?, ?, ?)`

	result, err := u.db.Exec(query, user.FirstName, user.LastName, user.Age,
		user.PhoneNumber, user.IsPhoneVerified)
	if err != nil {
		return 0, fmt.Errorf("INSERT INTO: %w", err)
	}

	newID, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("Get inserted ID: %w", err)
	}

	return int(newID), nil
}
