package main

import (
	"log"
	"net/http"

	"github.com/lifehou5e/homework/servergorilla/handlers"

	"github.com/gorilla/mux"
)

func main() {
	// init our router
	r := mux.NewRouter()

	r.HandleFunc("/users", handlers.CreateUser).Methods("POST")

	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err)
	}

}
