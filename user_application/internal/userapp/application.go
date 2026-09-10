package userapp

import (
	"fmt"
	"net/http"
	"userapp/internal/userapp/restapi"
)

func Run() {
	usersController := restapi.NewUsersController()
	router := restapi.NewRootRouter(usersController)

	s := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	fmt.Println("listening on port 8080")

	err := s.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
