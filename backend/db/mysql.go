package db

import (
	"database/sql"
	"log"
	"simvino/config"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/golang-migrate/migrate/source/file"
)

var Db *sql.DB

func initMySQL() {
	// Use root:dbpass@tcp(172.17.0.2)/hackernews, if you're using Windows.
	db, err := sql.Open("mysql", "root:"+config.SecretKeys.Mysql_pas+"@tcp(127.0.0.1:3306)/")
	if err != nil {
		log.Panic(err)
	}

	_, err = db.Exec("CREATE DATABASE IF NOT EXISTS simvino CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci; ")

	if err != nil {
		log.Panic(err)
	}

	db, err = sql.Open("mysql", "root:"+config.SecretKeys.Mysql_pas+"@tcp(127.0.0.1:3306)/simvino")
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
