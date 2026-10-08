package logic

import "time"

type User struct {
	ID         int
	FirstName  string
	MiddleName *string
	LastName   string
	Age        int
	Email      string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
