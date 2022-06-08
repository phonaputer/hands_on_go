package dal

import (
	"database/sql"
	"fmt"
	"github.com/go-sql-driver/mysql"
)

func NewMySQLDB() (*sql.DB, error) {
	dbConf := mysql.NewConfig()
	dbConf.User = "docker"
	dbConf.Passwd = "docker"
	dbConf.Addr = "localhost:3306"
	dbConf.DBName = "hands_on_go"

	connector, err := mysql.NewConnector(dbConf)
	if err != nil {
		return nil, fmt.Errorf("new connector: %w", err)
	}

	return sql.OpenDB(connector), nil
}
