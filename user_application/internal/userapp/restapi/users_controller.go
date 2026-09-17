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

func (c *UsersController) GetUser(w http.ResponseWriter, r *http.Request) {
	// 1. Get & validate input
	userIDStr := r.PathValue("user_id")

	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		fmt.Println("An error occurred: " + err.Error())
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("ERROR: User ID must be an integer!"))
		return
	}

	// 2. Get actual user data from logic layer
	user, err := c.crud.GetByID(r.Context(), userID)
	if err != nil {
		fmt.Println("An error occurred: " + err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// 3. Map result to HTTP response and send the reponse back
	response := c.mapper.User2UserResponse(user)

	jsonString, err := json.Marshal(response)
	if err != nil {
		fmt.Println("An error occurred: " + err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonString)
}
