package crypto

import (
	"fmt"
	"strconv"
	"time"

	"github.com/Hinterberger-Thomas/simvino/db"
)

func InsertNewCrypto(crypto []CryptoJson) error {
	tx, err := db.Db.Begin()
	if err != nil {
		return fmt.Errorf(err.Error()+" %#v %s", crypto, time.Now().String())
	}
	defer tx.Rollback()

	vals := []interface{}{}
	str := "INSERT INTO crypto (asset_id, asset_name, current_price, IS_CRYPTO) VALUES "
	j := 0
	for i := 0; i < len(crypto); i++ {
		str += fmt.Sprint("($" + strconv.FormatInt(int64(j+1), 10) + ", $" + strconv.FormatInt(int64(j+2), 10) + ", $" + strconv.FormatInt(int64(j+3), 10) + ", $" + strconv.FormatInt(int64(j+4), 10) + "),")
		vals = append(vals, crypto[i].AssetID, crypto[i].Name, crypto[i].PriceUsd, crypto[i].Is_crypto)
		j += 3
	}

	str = str[0 : len(str)-1]

	stmt, err := tx.Prepare(str)

	if err != nil {
		return fmt.Errorf(err.Error()+" %#v %s", crypto, time.Now().String())
	}
	defer stmt.Close()

	stmt.Exec(vals...)

	err = tx.Commit()
	if err != nil {
		return fmt.Errorf(err.Error()+" %#v %s", crypto, time.Now().String())
	}
	return nil
}

func GetCryptoById(crypto_Id uint32) (Crypto, error) {
	tx, err := db.Db.Begin()
	if err != nil {
		return Crypto{}, fmt.Errorf(err.Error()+" id: %d  time:%s", crypto_Id, time.Now().String())
	}
	defer tx.Rollback()

	stmt := "SELECT asset_id, asset_name, current_price FROM crypto WHERE crypto_id = $1"
	res := tx.QueryRow(stmt, crypto_Id)
	if err != nil {
		return Crypto{}, fmt.Errorf(err.Error()+" id: %d  time:%s", crypto_Id, time.Now().String())
	}

	crypto := Crypto{Asset_id: "", Asset_name: "", Crypto_id: 0}
	res.Scan(&crypto.Asset_id, &crypto.Asset_name, &crypto.Current_price)
	if err != nil {
		return Crypto{}, fmt.Errorf(err.Error()+" id: %d  time:%s", crypto_Id, time.Now().String())
	}
	return crypto, nil
}

func UpdateCryptoByAssetId(cryptoUpd CryptoUpd) error {
	tx, err := db.Db.Begin()
	if err != nil {
		return fmt.Errorf(err.Error()+" id: %#v  time:%s", cryptoUpd, time.Now().String())
	}
	defer tx.Rollback()

	stmt, err := tx.Prepare("UPDATE crypto SET current_price = $1  WHERE asset_id = $2")
	if err != nil {
		return fmt.Errorf(err.Error()+" id: %#v  time:%s", cryptoUpd, time.Now().String())
	}
	_, err = stmt.Exec(cryptoUpd.Current_price, cryptoUpd.Asset_id)

	err = tx.Commit()
	if err != nil {
		return fmt.Errorf(err.Error()+" id: %#v  time:%s", cryptoUpd, time.Now().String())
	}
	return nil
}

func GetCryptoByAssetId(asset_id string) (Crypto, error) {
	tx, err := db.Db.Begin()
	if err != nil {
		return Crypto{}, fmt.Errorf(err.Error()+" id: %s  time:%s", asset_id, time.Now().String())
	}
	defer tx.Rollback()

	stmt := "SELECT crypto_id, asset_name, current_price FROM crypto WHERE asset_id = $1"
	res := tx.QueryRow(stmt, asset_id)
	if err != nil {
		return Crypto{}, fmt.Errorf(err.Error()+" id: %s  time:%s", asset_id, time.Now().String())
	}

	crypto := Crypto{Asset_id: asset_id, Asset_name: "", Crypto_id: 0}
	res.Scan(&crypto.Crypto_id, &crypto.Asset_name, &crypto.Current_price)
	if err != nil {
		return Crypto{}, fmt.Errorf(err.Error()+" id: %s  time:%s", asset_id, time.Now().String())
	}
	return crypto, nil
}
