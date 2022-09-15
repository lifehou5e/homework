package main

import (
	"log"
	"net/http"
	"servergorilla/handlers"

	"github.com/gorilla/mux"
)

func main() {
	// init our router
	r := mux.NewRouter()

	// Some mock data
	// user = Users{Email: "ololo@gmail.com",
	// 	Password: "123456789abv",
	// 	FullName: "Oleh Naumov",
	// 	ID:       "123214"}

	// fmt.Printf("our firt user is: %+v\n", user)

	// route Handlers / Endpoints

	r.HandleFunc("/users", handlers.CreateUser).Methods("POST")

	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err)
	}

}
