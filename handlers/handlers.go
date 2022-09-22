package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/lifehou5e/homework/servergorilla/ent"
	"github.com/lifehou5e/homework/servergorilla/sqlconn"

	"github.com/google/uuid"
	_ "github.com/lib/pq"
)

var (
	user           ent.Users
	users          []ent.Users
	succesResponse = make(map[string]string, 0)
)

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

	user.ID = uuid.New()
	user.CreatedAt = time.Now()
	users = append(users, user)
	if len(responseErrors) != 0 {
		json.NewEncoder(w).Encode(responseErrors)
	} else {
		succesResponse["status"] = "Succes"
		json.NewEncoder(w).Encode(succesResponse)
		InsertIntoDB(&user)
	}
}

func InsertIntoDB(user *ent.Users) (int, error) {
	sqlStatement := `
INSERT INTO users (email, password, fullname)
VALUES ($1, $2, $3)
RETURNING id`
	id := 0
	err := sqlconn.DB.QueryRow(sqlStatement, user.Email, user.Password, user.FullName).Scan(&id)
	if err != nil {
		log.Println(err)
	}
	fmt.Println("New record ID is:", id)
	fmt.Println("Success!")
	return id, nil

}

func UniqueEmailCheck(email string) bool {
	sqlStatement := `SELECT EXISTS(SELECT 1 FROM users WHERE email = $1)`
	var exist bool
	err := sqlconn.DB.QueryRow(sqlStatement, user.Email).Scan(&exist)
	if err != nil {
		return false
	}

	return exist
}
