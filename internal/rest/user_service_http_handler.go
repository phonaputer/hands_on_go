package rest

import (
	"github.com/gorilla/mux"
	"net/http"
)

func NewUserServiceHTTPHandler(userController *UserController) http.Handler {
	r := mux.NewRouter()

	handleErr(r, "/users/{id}", userController.GetByID).Methods("GET")
	handleErr(r, "/users/{id}", userController.DeleteByID).Methods("DELETE")
	handleErr(r, "/users", userController.Create).Methods("POST")

	return r
}

func handleErr(r *mux.Router, path string, handlerFunc ErrHTTPHandler) *mux.Route {
	withErrResp := ErrResponseAdapter(handlerFunc)

	return r.Handle(path, withErrResp)
}
