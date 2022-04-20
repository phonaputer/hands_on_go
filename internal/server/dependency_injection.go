package server

import (
	"github.com/phonaputer/hands_on_go/internal/logic"
	"github.com/phonaputer/hands_on_go/internal/rest"
	"net/http"
)

type userServerApplication struct {
	userServerHTTPHandler http.Handler
}

func initializeUserServerApplication() (*userServerApplication, error) {

	userController := rest.NewUserController(
		&rest.UserValidatorStubImpl{},
		&logic.UserServiceStubImpl{},
	)

	userServerHTTPHandler := rest.NewUserServiceHTTPHandler(userController)

	return &userServerApplication{
		userServerHTTPHandler: userServerHTTPHandler,
	}, nil
}
