package db

import (
	"io/ioutil"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func createUser() {

	query, err := ioutil.ReadFile("db/sql/user.sql")

	if err != nil {
		log.Fatal(err)
	}

	_, err = Db.Exec(string(query))

	if err != nil {
		log.Fatal(err)
	}
}

func createBalance() {

	query, err := ioutil.ReadFile("db/sql/balance.sql")

	if err != nil {
		log.Fatal(err)
	}

	_, err = Db.Exec(string(query))

	if err != nil {
		log.Fatal(err)
	}
}
