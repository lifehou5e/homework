package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/lifehou5e/homework/servergorilla/ent"

	"github.com/google/uuid"
	_ "github.com/lib/pq"
)

var (
	user         ent.Users
	users        []ent.Users
	succesResone = make(map[string]string, 0)
)

func CreateUser(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "applictaion/json")

	responseErrors := make(map[string][]string)

	err := json.NewDecoder(req.Body).Decode(&user)
	if err != nil {
		return
	}
	if err := user.Validation(user); err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		responseErrors["error"] = append(responseErrors["error"], err.Error())
	}

	user.ID = uuid.New()
	user.CreatedAt = time.Now()
	users = append(users, user)
	// json.NewEncoder(w).Encode(user)
	if len(responseErrors) != 0 {
		json.NewEncoder(w).Encode(responseErrors)
	} else {
		succesResone["status"] = "Succes"
		json.NewEncoder(w).Encode(succesResone)
	}
	InsertIntoDB()
}

func InsertIntoDB() (int, error) {
	psqInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		ent.Host, ent.Port, ent.User, ent.Password, ent.Dbname)

	db, err := sql.Open("postgres", psqInfo)
	if err != nil {
		log.Fatal(err)
		return 0, err
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
		return 0, err
	}

	sqlStatement := `
INSERT INTO users (email, password, fullname, createdat, updatedat)
VALUES ($1, $2, $3, $4, $5)
RETURNING id`
	id := 0
	err = db.QueryRow(sqlStatement, "test@gmail.com", "Test Test", "passwordTest", time.Now(), time.Now()).Scan(&id)
	if err != nil {
		panic(err)
	}
	fmt.Println("New record ID is:", id)
	fmt.Println("Success!")
	return id, nil

}
