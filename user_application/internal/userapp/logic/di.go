package logic

type Component struct {
	UserCRUDImpl *UserCRUDImpl
}

type Dependencies struct {
	UsersRepository UserRepository
}

func NewComponent(d *Dependencies) *Component {
	var c Component

	c.UserCRUDImpl = newUserCRUDImpl(d.UsersRepository)

	return &c
}
