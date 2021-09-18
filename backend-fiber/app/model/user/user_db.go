package user

import (
	"github.com/Hinterberger-Thomas/simvino/auth"
	"github.com/Hinterberger-Thomas/simvino/db"
)

func InserUser(user User) (uint32, error) {
	tx, err := db.Db.Begin()
	if err != nil {
		return 0, err
	}
	defer tx.Rollback()

	stmt, err := tx.Prepare("INSERT INTO USERS (EMAIL, PASSWORD, ROLE) VALUES ($1, $2, $3) RETURNING user_id;")

	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	hash, err := auth.GeneratePassword(auth.DefaultConfig, user.Password)
	if err != nil {
		return 0, err
	}

	var userid uint32
	err = stmt.QueryRow(user.Email, hash, user.Role).Scan(&userid)
	if err != nil {
		return 0, err
	}
	err = tx.Commit()
	if err != nil {
		return 0, err
	}
	return userid, nil
}

func GetPasswordById(userId uint32) (string, error) {
	tx, err := db.Db.Begin()
	if err != nil {
		return "", err
	}
	defer tx.Rollback()

	stmt := "SELECT password FROM users WHERE user_id = $1"
	res := tx.QueryRow(stmt, userId)
	if err != nil {
		return "", err
	}

	var password string
	res.Scan(&password)
	return password, nil
}

func GetIdByEmail(email string) (uint32, error) {
	tx, err := db.Db.Begin()
	if err != nil {
		return 0, err
	}
	defer tx.Rollback()

	stmt := "SELECT user_id FROM users WHERE email = $1"
	res := tx.QueryRow(stmt, email)
	if err != nil {
		return 0, err
	}

	var userId uint32
	res.Scan(&userId)
	if err != nil {
		return 0, err
	}
	return userId, nil
}

func GetIdAndPassowrdByEmail(email string) (uint32, string, error) {
	tx, err := db.Db.Begin()
	if err != nil {
		return 0, "", err
	}
	defer tx.Rollback()

	stmt := "SELECT user_id, password FROM users WHERE email = $1"
	res := tx.QueryRow(stmt, email)
	if err != nil {
		return 0, "", err
	}

	var userId uint32
	var password string
	res.Scan(&userId, &password)
	if err != nil {
		return 0, "", err
	}
	return userId, password, nil
}
