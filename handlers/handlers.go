package handlers

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/lifehou5e/homework/servergorilla/ent"

	"github.com/google/uuid"
	_ "github.com/lib/pq"
)

var (
	user           ent.Users
	users          []ent.Users
	succesResponse = make(map[string]string, 0)
)

type Env struct {
	DB *sql.DB
}

func (env *Env) CreateUser(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "applictaion/json")

	responseErrors := make(map[string][]string)

	err := json.NewDecoder(req.Body).Decode(&user)
	if err != nil {
		return
	}
	if env.UniqueEmailCheck(user.Email) {
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
		env.InsertIntoDB(&user)
	}
}

func (env *Env) InsertIntoDB(user *ent.Users) error {
	sqlStatement := `
INSERT INTO users (first_name, last_name, email, password)
VALUES ($1, $2, $3, $4)
RETURNING id`
	err := env.DB.QueryRow(sqlStatement, user.FirstName, user.LastName, user.Email, user.Password)
	if err != nil {
		log.Println(err)
	}

	return nil
}

func (env *Env) UniqueEmailCheck(email string) bool {
	sqlStatement := `SELECT EXISTS(SELECT 1 FROM users WHERE email = $1)`

	var exist bool
	err := env.DB.QueryRow(sqlStatement, user.Email).Scan(&exist)
	if err != nil {
		return false
	}

	return exist
}
