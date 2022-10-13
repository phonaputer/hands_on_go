package dal

import (
	"database/sql"
	"fmt"
	"github.com/go-sql-driver/mysql"
)

type MySQLDBConf struct {
	User     string
	Password string
	Address  string
	DB       string
}

func NewMySQLDB(conf *MySQLDBConf) (*sql.DB, error) {
	driverConfig := mysql.NewConfig()

	driverConfig.User = conf.User
	driverConfig.Passwd = conf.Password
	driverConfig.Addr = conf.Address
	driverConfig.DBName = conf.DB

	conn, err := mysql.NewConnector(driverConfig)
	if err != nil {
		return nil, fmt.Errorf("mysql new connector: %w", err)
	}

	return sql.OpenDB(conn), nil
}
