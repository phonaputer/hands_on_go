package userapp

import (
	"hands_on_go/internal/presentation"
	"net/http"
)

type userApp struct {
	rootHandler http.Handler
}

func newUserApp() (*userApp, error) {
	var app userApp

	app.rootHandler = presentation.NewUserAppRootHandler()

	return &app, nil
}
