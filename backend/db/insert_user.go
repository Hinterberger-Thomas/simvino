package db

func (db *DB) InsertUser(user User) error {
	tx, err := db.client.Begin()

	if err != nil {
		return err
	}

	_, err = tx.Exec("INSERT INTO user (email, password) VALUES (?,?)", user.email, user.password)

	if err != nil {
		return err
	}

	return nil
}
