package sqlite

import "database/sql"

func UsersTableUp(db *sql.DB) error {
	query := `CREATE TABLE IF NOT EXISTS "USERS" (
		UserID INTEGER PRIMARY KEY AUTOINCREMENT,
		CreatedAt BIG INTEGER NOT NULL,
		EmailAddress TEXT UNIQUE NOT NULL,
		EncryptedPassword TEXT NOT NULL,
		FirstName TEXT NOT NULL,
		LastName TEXT NOT NULL,
		UpdatedAt BIG INTEGER NOT NULL,
		Username TEXT NOT NULL
	);`
	stmt, err := db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec()
	if err != nil {
		return err
	}
	return nil
}
