package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "simvino"
	password = "test"
	dbname   = "simvino"
)

var Db *sql.DB

func InitPostgres() {
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlconn)

	checkError(err)

	err = db.Ping()
	checkError(err)

	Db = db
	fmt.Println("Connected!")
}

func checkError(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
