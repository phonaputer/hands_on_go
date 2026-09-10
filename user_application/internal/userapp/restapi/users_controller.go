package restapi

import (
	"fmt"
	"net/http"
	"strconv"
)

type UsersController struct {
	// TODO
}

func NewUsersController() *UsersController {
	return &UsersController{}
}

func (c *UsersController) GetUser(w http.ResponseWriter, r *http.Request) {
	userIDStr := r.PathValue("user_id")

	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		fmt.Println("An error occurred: " + err.Error())
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("ERROR: User ID must be an integer!"))
		return
	}

	fmt.Printf("Get user: %d\n", userID)

	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("User data coming soon..."))
}
