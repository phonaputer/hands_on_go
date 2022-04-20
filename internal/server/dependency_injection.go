package server

import (
	"github.com/phonaputer/hands_on_go/internal/rest"
	"net/http"
)

type userServerApplication struct {
	userServerHTTPHandler http.Handler
}

func initializeUserServerApplication() (*userServerApplication, error) {
	userServerHTTPHandler := rest.NewUserServiceHTTPHandler()

	return &userServerApplication{
		userServerHTTPHandler: userServerHTTPHandler,
	}, nil
}
