package dal

import "github.com/jmoiron/sqlx"

type Component struct {
	MySQLUserRepository *MySQLUserRepository
}

type Dependencies struct {
	DB *sqlx.DB
}

func NewComponent(d *Dependencies) *Component {
	var c Component

	c.MySQLUserRepository = newMySQLUserRepository(d.DB)

	return &c
}
