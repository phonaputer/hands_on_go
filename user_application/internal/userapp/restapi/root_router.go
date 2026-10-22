package restapi

import "net/http"

func NewRootRouter(
	usersController *UsersController,
	errorResponseAdapter *errorResponseAdapter,
) http.Handler {
	m := muxer{
		router:               http.NewServeMux(),
		errorResponseAdapter: errorResponseAdapter,
	}

	m.HandleErrFunc("GET /users/{user_id}", usersController.GetUser)

	return m.router
}

type muxer struct {
	router               *http.ServeMux
	errorResponseAdapter *errorResponseAdapter
}

func (m *muxer) HandleErrFunc(path string, handler errHTTPHandlerFunc) {
	wrappedHandler := m.errorResponseAdapter.WithErrResp(handler)

	m.router.HandleFunc(path, wrappedHandler)
}
