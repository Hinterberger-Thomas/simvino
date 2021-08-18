package balances

import (
	"database/sql"
	"fmt"
	"log"
	"simvino/db"
)

type Balance struct {
	BalanceID int    `json:"balanceIds"`
	UserID    int    `json:"userId"`
	Currency  string `json:"currency"`
	Value     int    `json:"value"`
}

func (balance *Balance) InsertBalance() {
	statement, err := db.Db.Prepare("INSERT INTO balance (userID,currency,value) VALUES(?,?,?)")
	if err != nil {
		log.Fatal(err)
	}
	_, err = statement.Exec(balance.UserID, balance.Currency, balance.Value)
	if err != nil {
		fmt.Println(err)
	}
}

func GetUserByEmail(userID int) (Balance, error) {
	statement, err := db.Db.Prepare("select balanceID, currency, value from balance WHERE userID = ?")
	if err != nil {
		log.Fatal(err)
	}
	row := statement.QueryRow(userID)

	var balance Balance
	err = row.Scan(&balance.BalanceID, &balance.Currency, &balance.Value)
	if err != nil {
		if err != sql.ErrNoRows {
			log.Print(err)
		}
		return balance, err
	}

	return balance, nil
}
