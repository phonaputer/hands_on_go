package userapp

import (
	"fmt"
	"net/http"
	"userapp/internal/userapp/dal"
	"userapp/internal/userapp/logic"
	"userapp/internal/userapp/restapi"
)

// curl localhost:8080/users/123
// -> should return the dummy user JSON

func Run() {
	dataAccessLayer := dal.NewComponent(&dal.Dependencies{})

	logicLayer := logic.NewComponent(&logic.Dependencies{
		UsersRepository: dataAccessLayer.MySQLUserRepository,
	})

	restAPI := restapi.NewComponent(&restapi.Dependencies{
		UserCRUD: logicLayer.UserCRUDImpl,
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
