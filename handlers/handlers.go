package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/lifehou5e/homework/servergorilla/ent"

	"github.com/google/uuid"
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
}
