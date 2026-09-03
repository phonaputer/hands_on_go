package userapp

import (
	"fmt"
	"net/http"
)

func Run() {
	s := &http.Server{
		Addr:    ":8080",
		Handler: http.HandlerFunc(handleRequest),
	}

	err := s.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	fmt.Println("I was called!")

	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Thanks for calling!\n"))
}
