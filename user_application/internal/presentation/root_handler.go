package presentation

import (
	"github.com/gorilla/mux"
	"net/http"
)

// NewUserAppRootHandler initializes the top-level HTTP request router for User Application.
func NewUserAppRootHandler(userController *UserController) http.Handler {
	r := mux.NewRouter()

	r.HandleFunc("/users", userController.GetUserByID).Methods("GET")
	r.HandleFunc("/users", userController.CreateUser).Methods("POST")

	return r
}
