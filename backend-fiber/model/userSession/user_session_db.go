package userSession

import (
	"fmt"

	"github.com/Hinterberger-Thomas/simvino/db"
)

func InsertSession(userSessoinIns Users_Sessions_Ins) error {
	tx, err := db.Db.Begin()
	if err != nil {
		return fmt.Errorf(err.Error()+" %#v", userSessoinIns)
	}
	defer tx.Rollback()

	stmt, err := tx.Prepare("INSERT INTO users_sessions (fk_user_id, session_id) VALUES ($1, $2)")

	if err != nil {
		return fmt.Errorf(err.Error()+" %#v", userSessoinIns)
	}
	defer stmt.Close()

	_, err = stmt.Exec(userSessoinIns.Fk_user_id, userSessoinIns.SESSION_ID)
	if err != nil {
		return fmt.Errorf(err.Error()+" %#v", userSessoinIns)
	}

	err = tx.Commit()
	if err != nil {
		return fmt.Errorf(err.Error()+" %#v", userSessoinIns)
	}
	return nil
}
