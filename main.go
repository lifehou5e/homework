package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"github.com/lifehou5e/homework/servergorilla/handlers"
	"github.com/lifehou5e/homework/servergorilla/sqlconn"
)

func main() {
	// init our router
	r := mux.NewRouter()

	var err error

	db, err := sql.Open("postgres", fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		sqlconn.Host, sqlconn.Port, sqlconn.User, sqlconn.Password, sqlconn.DBname))
	if err != nil {
		log.Println(err)
	}
	defer db.Close()

	env := &handlers.Env{DB: db}

	err = db.Ping()
	if err != nil {
		log.Println(err)
	}

	r.HandleFunc("/users", env.CreateUser).Methods("POST")

	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err)
	}

}
