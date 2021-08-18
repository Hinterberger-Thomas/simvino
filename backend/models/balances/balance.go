package balances

import (
	"fmt"
	"log"
	"simvino/db"
	"simvino/graph/model"
)

type Balance struct {
	BalanceID int    `json:"balanceIds"`
	UserID    int    `json:"userId"`
	Currency  string `json:"currency"`
	Value     int    `json:"value"`
}

func InsertBalance(balance *Balance) {
	statement, err := db.Db.Prepare("INSERT INTO balance (userID,currency,value) VALUES(?,?,?)")
	if err != nil {
		log.Fatal(err)
	}
	_, err = statement.Exec(balance.UserID, balance.Currency, balance.Value)
	if err != nil {
		fmt.Println(err)
	}
}

func GetTransactionByUserID(userID int) []*model.Transaction {
	statement, err := db.Db.Prepare("select currency, value from balance WHERE userID = ?")
	if err != nil {
		log.Fatal(err)
	}
	results, err := statement.Query(userID)

	if err != nil {
		fmt.Println(err)
	}

	var transactions []*model.Transaction

	for results.Next() {
		var tranaciton model.Transaction

		err = results.Scan(&tranaciton.Currency, &tranaciton.Value)
		if err != nil {
			panic(err.Error())
		}
		transactions = append(transactions, &tranaciton)
	}

	return transactions
}
