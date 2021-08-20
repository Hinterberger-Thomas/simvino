package users

import (
	"crypto/subtle"
	"database/sql"
	"fmt"
	"simvino/db"

	"golang.org/x/crypto/argon2"

	"log"
)

type User struct {
	UserID   int    `json:"userId"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (user *User) InsertUser() error {
	statement, err := db.Db.Prepare("INSERT INTO user (email,Password) VALUES(?,?)")

	if err != nil {
		return err
	}

	salt := generatePassword()
	insertSalt(user.Email, salt)
	hashedPassword := HashPassword(user.Password, salt)
	_, err = statement.Exec(user.Email, hashedPassword)
	if err != nil {
		return &DuplicateEmail{}
	}
	return err
}

//HashPassword hashes given password
func HashPassword(password string, salt string) string {
	hash := argon2.IDKey([]byte(password), []byte(salt), 1, 128*1024, 4, 32)
	return string(hash)
}

//CheckPassword hash compares raw password with it's hashed values
func CheckPasswordHash(password, hash string, salt string) bool {
	passwordHashed := argon2.IDKey([]byte(password), []byte(salt), 2, 128*1024, 4, 32)
	return subtle.ConstantTimeCompare([]byte(passwordHashed), []byte(hash)) == 1
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
	salt, err := getSalt(user.Email)
	if err != nil {
		fmt.Println(err)
	}
	return CheckPasswordHash(user.Password, hashedPassword, salt)
}

func insertSalt(email string, salt string) error {
	err := db.Client.Set(email, salt, 0).Err()
	// if there has been an error setting the value
	// handle the error
	if err != nil {
		return err
	}
	return nil
}

func getSalt(email string) (string, error) {
	val, err := db.Client.Get(email).Result()
	if err != nil {
		return "", err
	}

	return val, nil
}
