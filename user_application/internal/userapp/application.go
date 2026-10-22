package userapp

import (
	"fmt"
	"net/http"
	"userapp/internal/userapp/dal"
	"userapp/internal/userapp/logic"
	"userapp/internal/userapp/restapi"

	_ "github.com/go-sql-driver/mysql"

	"github.com/jmoiron/sqlx"
)

func Run() {
	db, err := sqlx.Open("mysql", "dockeruser:dockerpass@tcp(localhost:3306)/hands_on_go?parseTime=true")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	dataAccessLayer := dal.NewComponent(&dal.Dependencies{
		DB: db,
	})

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

	err = s.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
