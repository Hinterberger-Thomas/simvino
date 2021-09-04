package main

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	host := "127.0.0.1"
	port := "5432"
	user := "postgres"
	password := "example"
	dbname := "simvino"

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatalf("Tidak Konek DB Errornya : %s", err)
	}
	defer db.Close()

	query, err := ioutil.ReadFile("user.sql")

	if err != nil {
		fmt.Println(err)
	}
	_, err = db.Exec(string(query))

	if err != nil {
		log.Fatalf("Tidak Konek DB Errornya : %s", err)
	}

	statement := `INSERT INTO "user" ("email","Password") VALUES($1,$2)`

	if err != nil {
		fmt.Println(err)
	}

	_, err = db.Exec(statement, "jojo", "\xb3lϱJ\xbf\xa8\xfd䨰&]\xe1*\x99*\xfc\xa9\f\xd3\xeeTS\x11ʮ\xa8>\x98&`")

	if err != nil {
		fmt.Println(err)
	}

}
