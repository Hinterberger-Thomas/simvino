package users

import (
	"database/sql"
	"fmt"
	"simvino/db"

	"golang.org/x/crypto/bcrypt"

	"log"
)

type User struct {
	UserID   int    `json:"userId"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (user *User) InsertUser() {
	statement, err := db.Db.Prepare("INSERT INTO user (email,Password) VALUES(?,?)")
	print(statement)
	if err != nil {
		log.Fatal(err)
	}
	hashedPassword, _ := HashPassword(user.Password)
	_, err = statement.Exec(user.Email, hashedPassword)
	if err != nil {
		fmt.Println(err)
	}
}

//HashPassword hashes given password
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

//CheckPassword hash compares raw password with it's hashed values
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func GetUserByEmail(email string) (User, error) {
	statement, err := db.Db.Prepare("select userID, password,email from user WHERE email = ?")
	if err != nil {
		log.Fatal(err)
	}
	row := statement.QueryRow(email)

	var user User
	err = row.Scan(&user.UserID, &user.Password, &user.Email)
	if err != nil {
		if err != sql.ErrNoRows {
			log.Print(err)
		}
		return User{}, err
	}

	return user, nil
}

func (user *User) Authenticate() bool {
	statement, err := db.Db.Prepare("select password from user WHERE email = ?")
	if err != nil {
		log.Fatal(err)
	}
	row := statement.QueryRow(user.Email)

	var hashedPassword string
	err = row.Scan(&hashedPassword)
	if err != nil {
		if err == sql.ErrNoRows {
			return false
		} else {
			log.Fatal(err)
		}
	}

	return CheckPasswordHash(user.Password, hashedPassword)
}
