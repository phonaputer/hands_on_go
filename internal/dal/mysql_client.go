package dal

import (
	"database/sql"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"github.com/phonaputer/hands_on_go/internal/config"
)

func NewMySQLDB(conf config.MySQL) (*sql.DB, error) {
	dbConf := mysql.NewConfig()
	dbConf.User = conf.Username
	dbConf.Passwd = conf.Password
	dbConf.Addr = conf.Address
	dbConf.DBName = conf.DB

	connector, err := mysql.NewConnector(dbConf)
	if err != nil {
		return nil, fmt.Errorf("new connector: %w", err)
	}

	return sql.OpenDB(connector), nil
}
