package cryptoHistory

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/Hinterberger-Thomas/simvino/db"
	"go.mongodb.org/mongo-driver/bson"
)

// func InsertCryptoHis(cryptoHis []crypto.CryptoJson) error {
// 	tx, err := db.Db.Begin()
// 	if err != nil {
// 		return fmt.Errorf(err.Error()+" %#v", cryptoHis)
// 	}
// 	defer tx.Rollback()

// 	vals := []interface{}{}
// 	str := "INSERT INTO crypto_history (price, fk_asset_id, date) VALUES "
// 	j := 0
// 	for i := 0; i < len(cryptoHis); i++ {
// 		str += fmt.Sprint("($" + strconv.FormatInt(int64(j+1), 10) + ", $" + strconv.FormatInt(int64(j+2), 10) + ", $" + strconv.FormatInt(int64(j+3), 10) + "),")
// 		vals = append(vals, cryptoHis[i].PriceUsd, cryptoHis[i].AssetID, time.Now())
// 		j += 3
// 	}

// 	str = str[0 : len(str)-1]

// 	stmt, err := tx.Prepare(str)

// 	if err != nil {
// 		return fmt.Errorf(err.Error()+" %#v", cryptoHis)
// 	}
// 	defer stmt.Close()

// 	_, err = stmt.Exec(vals...)
// 	if err != nil {
// 		return fmt.Errorf(err.Error()+" %#v", cryptoHis)
// 	}

// 	err = tx.Commit()
// 	if err != nil {
// 		return fmt.Errorf(err.Error()+" %#v", cryptoHis)
// 	}
// 	return nil
// }

func InsertHis(cryptoHis []interface{}) error {
	coll := db.MoClient.Database("history").Collection("crypto")

	_, e := coll.InsertMany(context.TODO(), cryptoHis)
	return e
}

func Get(from time.Time, to time.Time, asset_id string) error {
	coll := db.MoClient.Database("history").Collection("crypto")

	fmt.Println(from)
	fmt.Println(to)

	cur, err := coll.Find(context.TODO(), bson.M{"time": bson.M{
		"$gte": from.Unix(),
		"$lt":  to.Unix(),
	}, "asset_id": strings.ToUpper(asset_id)})
	if err != nil {
		return err
	}
	var epi []CryptoHistroyIns
	if err = cur.All(context.TODO(), &epi); err != nil {
		log.Fatal(err)
	}
	return nil
}
