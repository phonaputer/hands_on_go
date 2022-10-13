package userapp

import (
	"database/sql"
	"fmt"
	"github.com/sirupsen/logrus"
	"hands_on_go/internal/dal"
	"hands_on_go/internal/logic"
	"hands_on_go/internal/presentation"
	"net/http"
)

type userApp struct {
	mysqlDB *sql.DB

	rootHandler http.Handler
}

func newUserApp(conf *userAppConfig) (*userApp, error) {
	var app userApp

	db, err := dal.NewMySQLDB(&dal.MySQLDBConf{
		User:     conf.MySQL.User,
		Password: conf.MySQL.Password,
		Address:  conf.MySQL.Address,
		DB:       conf.MySQL.Database,
	})
	if err != nil {
		return nil, fmt.Errorf("open mysql pool: %w", err)
	}

	app.mysqlDB = db

	userRepo := dal.NewUserRepository(app.mysqlDB)

	userService := logic.NewUserService(userRepo)

	userController := presentation.NewUserController(
		&presentation.UserValidatorImpl{},
		userService,
	)

	app.rootHandler = presentation.NewUserAppRootHandler(userController)

	return &app, nil
}

func closeUserApp(app *userApp) {
	err := app.mysqlDB.Close()
	if err != nil {
		logrus.WithError(err).Error("failed to close MySQL client")
	}
}
