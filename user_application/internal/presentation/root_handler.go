package presentation

import "net/http"

// NewUserAppRootHandler initializes the top-level HTTP request router for User Application.
func NewUserAppRootHandler() http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Write([]byte("Hello world!"))
	})

}
