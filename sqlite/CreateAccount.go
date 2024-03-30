package sqlite

import (
	"database/sql"
	"log"
	"matthewhope/atm-system-go/models"
)

func CreateAccount(db *sql.DB, a models.Account) (models.Account, error) {
	ret := a
	query := `INSERT INTO "ACCOUNTS" (
		CreatedAt,
		AccountNumber,
		Balance,
		CreatedAt,
		UpdatedAt,
		UserID
	) VALUES (?, ?, ?, ?, ?, ?);`
	stmt, err := db.Prepare(query)
	if err != nil {
		log.Println(err.Error())
		return ret, err
	}
	defer stmt.Close()
	res, err := stmt.Exec(a.AccountNumber, a.Balance, a.CreatedAt, a.UpdatedAt, a.UserID)
	if err != nil {
		return ret, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return ret, err
	}
	ret.UserID = int(id)
	return ret, nil
}
