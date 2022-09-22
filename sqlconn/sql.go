package sqlconn

import (
	"database/sql"
	"fmt"
	"log"
)

const (
	Host     = "localhost"
	Port     = 5432
	User     = "lifehou5e"
	Password = "new_password"
	Dbname   = "users"
)

var DB *sql.DB

func OpenDBConnection() (db *sql.DB, err error) {
	psqInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		Host, Port, User, Password, Dbname)

	db, err = sql.Open("postgres", psqInfo)
	if err != nil {
		log.Println(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Println(err)
	}

	return db, nil
}
