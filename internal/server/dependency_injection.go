package server

import (
	"fmt"
	"github.com/phonaputer/hands_on_go/internal/config"
	"github.com/phonaputer/hands_on_go/internal/dal"
	"github.com/phonaputer/hands_on_go/internal/logic"
	"github.com/phonaputer/hands_on_go/internal/rest"
	"net/http"
)

type userServerApplication struct {
	userServerHTTPHandler http.Handler
}

func initializeUserServerApplication(conf *config.HOGConfig) (*userServerApplication, error) {

	db, err := dal.NewMySQLDB(conf.MySQL)
	if err != nil {
		return nil, fmt.Errorf("new MySQL DB: %w", err)
	}

	userRepository := dal.NewUserRepositoryMySQLImpl(db)

	userService := logic.NewUserServiceImpl(userRepository)

	userController := rest.NewUserController(
		&rest.UserValidatorImpl{},
		userService,
	)

	userServerHTTPHandler := rest.NewUserServiceHTTPHandler(userController)

	return &userServerApplication{
		userServerHTTPHandler: userServerHTTPHandler,
	}, nil
}
