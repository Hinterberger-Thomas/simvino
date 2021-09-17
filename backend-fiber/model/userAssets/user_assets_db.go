package userAssets

import (
	"github.com/Hinterberger-Thomas/simvino/db"
)

func InserUserAsset(userAsset User_assets_ins) error {
	tx, err := db.Db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	stmt, err := tx.Prepare("INSERT INTO users_assets (fk_crypto_id, fk_user_id, amount) VALUES  ($1, $2, $3);")

	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(userAsset.Crypto_id, userAsset.Fk_user_id, userAsset.Amount)
	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}
	return err
}
