package logic

type Component struct {
	UserCRUDImpl *UserCRUDImpl
}

type Dependencies struct {
	UserRepository UserRepository
}

func NewComponent(d *Dependencies) *Component {
	var c Component

	c.UserCRUDImpl = newUserCRUDImpl(d.UserRepository)

	return &c
}
