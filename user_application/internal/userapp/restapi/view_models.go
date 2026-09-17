package restapi

import "time"

type userResponse struct {
	ID         int       `json:"id"`
	FirstName  string    `json:"first_name"`
	MiddleName *string   `json:"middle_name,omitempty"`
	LastName   string    `json:"last_name"`
	Age        int       `json:"age"`
	Email      string    `json:"email"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
