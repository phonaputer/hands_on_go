package userapp

import (
	"hands_on_go/internal/logic"
	"hands_on_go/internal/presentation"
	"net/http"
)

type userApp struct {
	rootHandler http.Handler
}

func newUserApp() (*userApp, error) {
	var app userApp

	userController := presentation.NewUserController(
		&presentation.UserValidatorImpl{},
		&logic.UserServiceStubImpl{},
	)

	app.rootHandler = presentation.NewUserAppRootHandler(userController)

	return &app, nil
}
