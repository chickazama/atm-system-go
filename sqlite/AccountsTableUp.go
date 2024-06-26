package sqlite

import "database/sql"

func AccountsTableUp(db *sql.DB) error {
	query := `CREATE TABLE IF NOT EXISTS "ACCOUNTS" (
		AccountID INTEGER PRIMARY KEY AUTOINCREMENT,
		AccountNumber TEXT UNIQUE NOT NULL,
		Balance BIG INTEGER NOT NULL,
		CreatedAt BIG INTEGER NOT NULL,
		UpdatedAt BIG INTEGER NOT NULL,
		UserID INTEGER NOT NULL
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
