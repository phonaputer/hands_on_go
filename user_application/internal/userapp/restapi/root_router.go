package restapi

import "net/http"

func NewRouter(usersController *UsersController) http.Handler {
	router := http.NewServeMux()

	router.HandleFunc("GET /users/{user_id}", usersController.GetUser)

	return router
}
