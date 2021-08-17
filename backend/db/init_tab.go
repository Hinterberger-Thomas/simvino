package db

import (
	"database/sql"
	"io/ioutil"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type DB struct {
	client *sql.DB
}

func OpenCon() *sql.DB {
	db, err := sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/simvino")

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	return db

}

func (db *DB) createUser() {

	query, err := ioutil.ReadFile("db/sql/user.sql")

	if err != nil {
		log.Fatal(err)
	}

	_, err = db.client.Exec(string(query))

	if err != nil {
		log.Fatal(err)
	}

}
