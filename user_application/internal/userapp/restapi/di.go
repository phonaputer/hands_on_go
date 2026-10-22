package restapi

import (
	"net/http"
	"userapp/internal/userapp/logic"
)

type Component struct {
	RootRouter http.Handler

	errorResponseAdapter *errorResponseAdapter

	usersController *UsersController
	jsonUsersMapper jsonUsersMapper
}

type Dependencies struct {
	UserCRUD logic.UserCRUD
}

func NewComponent(d *Dependencies) *Component {
	var c Component

	c.jsonUsersMapper = newJSONUsersMapper()

	c.usersController = NewUsersController(
		d.UserCRUD,
		c.jsonUsersMapper,
	)

	c.errorResponseAdapter = newErrorResponseAdapter()

	c.RootRouter = NewRootRouter(c.usersController, c.errorResponseAdapter)

	return &c
}
