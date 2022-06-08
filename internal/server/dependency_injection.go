package server

import (
	"github.com/phonaputer/hands_on_go/internal/dal"
	"github.com/phonaputer/hands_on_go/internal/logic"
	"github.com/phonaputer/hands_on_go/internal/rest"
	"net/http"
)

type userServerApplication struct {
	userServerHTTPHandler http.Handler
}

func initializeUserServerApplication() (*userServerApplication, error) {

	userService := logic.NewUserServiceImpl(&dal.UserRepositoryStub{})

	userController := rest.NewUserController(
		&rest.UserValidatorImpl{},
		userService,
	)

	userServerHTTPHandler := rest.NewUserServiceHTTPHandler(userController)

	return &userServerApplication{
		userServerHTTPHandler: userServerHTTPHandler,
	}, nil
}
