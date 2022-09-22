package main

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"io"
	"log"
	"net/http"
	"os"
	"unicode"
)

const keyServerAddr = "serverAddr"

func getUser(w http.ResponseWriter, r *http.Request) {
	//////////////Inspecting a Request's Query String/////////////////////
	userid := r.URL.Query().Get("id")

	if userid == "" {
		w.Header().Set("x-missing-field", "userID")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	for _, r := range userid {
		if !unicode.IsNumber(r) {
			w.Header().Set("x-missing-field", "userID")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}

	//////////////Database/////////////////////
	db, err := sql.Open("mysql",
		"dockeruser:dockerpass@tcp(127.0.0.1:3306)/hands_on_go")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	var (
		id         int
		first_name string
		last_name  string
	)
	rows, err := db.Query("select id, first_name, last_name from users where id = ?", userid)

	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&id, &first_name, &last_name)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(id, first_name, last_name)
	}

	if id == 0 {
		w.Header().Set("x-missing-field", "non existance user")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	//////////////Mashalling json data (encoding into JSON)/////////////////////
	type User struct {
		Id        int    `json:"id"`
		Firstname string `json:"first_name"`
		Lastname  string `json:"last_name"`
	}

	user := &User{
		Id:        id,
		Firstname: first_name,
		Lastname:  last_name,
	}

	data, _ := json.Marshal(user)
	io.WriteString(w, string(data))
}

func main() {
	http.HandleFunc("/get-user", getUser)

	err := http.ListenAndServe(":8080", nil)

	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
}
