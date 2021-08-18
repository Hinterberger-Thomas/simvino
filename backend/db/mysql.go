package db

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/golang-migrate/migrate/source/file"
)

var Db *sql.DB

func InitDB() {
	// Use root:dbpass@tcp(172.17.0.2)/hackernews, if you're using Windows.
	db, err := sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/simvino")
	if err != nil {
		log.Panic(err)
	}

	if err = db.Ping(); err != nil {
		log.Panic(err)
	}
	Db = db

	createUser()
	createBalance()
}

// func Migrate() {
// 	if err := Db.Ping(); err != nil {
// 		log.Fatal(err)
// 	}
// 	driver, _ := mysql.WithInstance(Db, &mysql.Config{})
// 	m, err := migrate.NewWithDatabaseInstance(
// 		"file://internal/pkg/db/migrations/mysql",
// 		"mysql",
// 		driver,
// 	)

// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
// 		log.Fatal(err)
// 	}

// }
