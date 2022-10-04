package presentation

import (
	"github.com/gorilla/mux"
	"net/http"
)

// NewUserAppRootHandler initializes the top-level HTTP request router for User Application.
func NewUserAppRootHandler(userController *UserController) http.Handler {
	r := mux.NewRouter()

	//r.HandleFunc("/users", userController.GetUserByID).Methods("GET")
	handleErr(r, "/users", userController.GetUserByID).Methods("GET")
	handleErr(r, "/users", userController.CreateUser).Methods("POST")
	handleErr(r, "/users", userController.DeleteUserByID).Methods("DELETE")

	return r
}

func handleErr(r *mux.Router, path string, handler ErrHTTPHandler) *mux.Route {
	return r.Handle(path, WithErrResponse(handler))
}
