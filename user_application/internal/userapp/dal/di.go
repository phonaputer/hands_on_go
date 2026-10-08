package dal

type Component struct {
	MySQLUserRepository *MySQLUserRepository
}

type Dependencies struct {
}

func NewComponent(d *Dependencies) *Component {
	var c Component

	c.MySQLUserRepository = newMySQLUserRepository()

	return &c
}
