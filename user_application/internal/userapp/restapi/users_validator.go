package restapi

import (
	"encoding/json"
	"net/http"
	"userapp/internal/userapp/logic"
)

type usersValidator interface {
	Create(r *http.Request) (*logic.NewUser, error)
}

type jsonUsersValidator struct{}

// TODO implement this

func (j jsonUsersValidator) Create(r *http.Request) (*logic.NewUser, error) {
	var reqJSON struct {
	}

	err := json.NewDecoder(r.Body).Decode(&reqJSON)

	return &logic.NewUser{}, nil
}
