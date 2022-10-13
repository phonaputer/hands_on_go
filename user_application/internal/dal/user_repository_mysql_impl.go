package dal

import (
	"database/sql"
	"errors"
	"fmt"
	"hands_on_go/internal/logic"
	"hands_on_go/internal/uaerr"
)

type userRepositoryMySQLImpl struct {
	db *sql.DB
}

func NewUserRepository(
	db *sql.DB,
) logic.UserRepository {
	return &userRepositoryMySQLImpl{
		db: db,
	}
}

func (u *userRepositoryMySQLImpl) CreateUser(user *logic.User) (int, error) {
	const query = `INSERT INTO users
							(first_name, last_name, age, phone_number, phone_verification_status)
						VALUES 
						    (?, ?, ?, ?, ?)`

	res, err := u.db.Exec(query, user.FirstName, user.LastName, user.Age, user.PhoneNumber, user.PhoneVerified)
	if err != nil {
		return 0, fmt.Errorf("INSERT user: %w", err)
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("get inserted user ID: %w", err)
	}

	return int(id), nil
}

func (u *userRepositoryMySQLImpl) GetUser(id int) (*logic.User, error) {
	const query = `SELECT 
						first_name, last_name, age, phone_number, phone_verification_status
					FROM
						users
					WHERE
						id=?`

	row := u.db.QueryRow(query, id)

	var result logic.User

	err := row.Scan(&result.FirstName, &result.LastName, &result.Age, &result.PhoneNumber, &result.PhoneVerified)
	if errors.Is(err, sql.ErrNoRows) {
		err = uaerr.SetType(err, uaerr.TypeNotFound)
		return nil, uaerr.SetUserMsg(err, "user not found")
	}
	if err != nil {
		return nil, fmt.Errorf("SELECT user: %w", err)
	}

	return &result, nil
}

func (u *userRepositoryMySQLImpl) CheckUserExists(id int) (bool, error) {
	const query = `SELECT count(*) FROM users WHERE id=?`

	row := u.db.QueryRow(query, id)

	var count int

	err := row.Scan(&count)
	if errors.Is(err, sql.ErrNoRows) {
		return false, nil
	}
	if err != nil {
		return false, fmt.Errorf("SELECT count(*) from users: %w", err)
	}

	return count > 0, nil
}

func (u *userRepositoryMySQLImpl) DeleteUser(id int) error {
	const query = `DELETE FROM users WHERE id=?`

	_, err := u.db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("DELETE user: %w", err)
	}

	return nil
}
