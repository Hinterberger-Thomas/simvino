package db

import (
	"log"
)

type User struct {
	userId   int
	email    string
	password string
}

func (db *DB) GetUser(userid int) User {
	var us User
	tx, err := db.client.Begin()

	if err != nil {
		log.Fatal(err)
	}

	row := tx.QueryRow("SELECT userid, email, password from user WHERE userid = ?", userid)
	err = row.Scan(&us)

	if err != nil {
		log.Fatal(err)
	}

	return us
}
