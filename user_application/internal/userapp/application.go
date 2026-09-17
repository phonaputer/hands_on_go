package userapp

import (
	"fmt"
	"net/http"
	"userapp/internal/userapp/logic"
	"userapp/internal/userapp/restapi"
)

func Run() {
	restAPI := restapi.NewComponent(&restapi.Dependencies{
		UserCRUD: logic.NewDummyUserCRUD(),
	})

	s := &http.Server{
		Addr:    ":8080",
		Handler: restAPI.RootRouter,
	}

	fmt.Println("listening on port 8080")

	err := s.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
