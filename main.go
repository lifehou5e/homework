package main

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/lifehou5e/homework/servergorilla/validation"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

const (
	Host     = "localhost"
	Port     = 5432
	User     = "lifehou5e"
	Password = "new_password"
	DBname   = "users"
)

var (
	user           Users
	users          []Users
	succesResponse = make(map[string]string, 0)
	DB             *sql.DB
)

type Users struct {
	ID        uuid.UUID
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func (u *Users) Validation(user Users) error {
	if validation.IsASCII(u.Password) == false {
		return errors.New("password must contain only ASCII symbols")
	}
	if len(u.Password) < 8 {
		return errors.New("password should have at least 8 characters and be less than 256 symbols")
	}
	if len(u.Password) >= 256 {
		return errors.New("password should have less than 256 characters")
	}
	if len(u.Email) >= 256 {
		return errors.New("email should have less than 256 characters")
	}
	if !validation.ContainDog(u.Email) {
		return errors.New("invalid email")
	}

	return nil
}

func main() {
	// init our router
	r := mux.NewRouter()

	var err error

	DB, err = sql.Open("postgres", fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		Host, Port, User, Password, DBname))
	if err != nil {
		log.Println(err)
	}
	defer DB.Close()

	err = DB.Ping()
	if err != nil {
		log.Println(err)
	}

	r.HandleFunc("/users", CreateUser).Methods("POST")

	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err)
	}

}

func CreateUser(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "applictaion/json")

	responseErrors := make(map[string][]string)

	err := json.NewDecoder(req.Body).Decode(&user)
	if err != nil {
		return
	}
	if UniqueEmailCheck(user.Email) {
		w.WriteHeader(http.StatusConflict)
		responseErrors["error"] = append(responseErrors["error"], "this email already has been taken")
	}
	if err := user.Validation(user); err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		responseErrors["error"] = append(responseErrors["error"], err.Error())
	}

	users = append(users, user)
	if len(responseErrors) != 0 {
		json.NewEncoder(w).Encode(responseErrors)
	} else {
		succesResponse["status"] = "Succes"
		json.NewEncoder(w).Encode(succesResponse)
		InsertIntoDB(&user)
	}
}

func InsertIntoDB(user *Users) error {
	sqlStatement := `
INSERT INTO users (first_name, last_name, email, password)
VALUES ($1, $2, $3, $4)
RETURNING id`
	err := DB.QueryRow(sqlStatement, user.FirstName, user.LastName, user.Email, user.Password)
	if err != nil {
		log.Println(err)
	}

	return nil
}

func UniqueEmailCheck(email string) bool {
	sqlStatement := `SELECT EXISTS(SELECT 1 FROM users WHERE email = $1)`
	var exist bool
	err := DB.QueryRow(sqlStatement, user.Email).Scan(&exist)
	if err != nil {
		return false
	}

	return exist
}
