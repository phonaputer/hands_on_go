package rest

import (
	"github.com/gorilla/mux"
	"net/http"
)

func NewUserServiceHTTPHandler(userController *UserController) http.Handler {
	r := mux.NewRouter()

	r.HandleFunc("/users/{id}", userController.GetByID).Methods("GET")
	r.HandleFunc("/users/{id}", userController.DeleteByID).Methods("DELETE")
	r.HandleFunc("/users", userController.Create).Methods("POST")

	return r
}
