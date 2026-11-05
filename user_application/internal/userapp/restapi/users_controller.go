package restapi

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"userapp/internal/userapp/logic"
)

type UsersController struct {
	crud   logic.UserCRUD
	mapper usersMapper
}

func NewUsersController(
	crud logic.UserCRUD,
	mapper usersMapper,
) *UsersController {
	return &UsersController{
		crud:   crud,
		mapper: mapper,
	}
}

func (c *UsersController) GetUser(w http.ResponseWriter, r *http.Request) error {
	// 1. Get & validate input
	userIDStr := r.PathValue("user_id")

	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		err = fmt.Errorf("%w: convert id to integer: %w", logic.ErrInvalidInput, err)
		return logic.SetEndUserMsg(err, "User ID must be an integer!")
	}

	// 2. Get actual user data from logic layer
	user, err := c.crud.GetByID(r.Context(), userID)
	if err != nil {
		return fmt.Errorf("get by id: %w", err)
	}

	// 3. Map result to HTTP response and send the reponse back
	response := c.mapper.User2UserResponse(user)

	jsonString, err := json.Marshal(response)
	if err != nil {
		return fmt.Errorf("json marshal response: %w", err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonString)

	return nil
}
