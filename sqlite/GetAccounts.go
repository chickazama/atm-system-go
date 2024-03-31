package sqlite

import (
	"database/sql"
	"matthewhope/atm-system-go/models"
)

func GetAccounts(db *sql.DB) ([]models.Account, error) {
	var ret []models.Account
	query := `SELECT * FROM "USERS";`
	rows, err := db.Query(query)
	if err != nil {
		return ret, nil
	}
	for rows.Next() {
		var a models.Account
		err = rows.Scan(&a.AccountID, &a.AccountNumber, &a.Balance, &a.CreatedAt, &a.UpdatedAt, &a.UserID)
		if err != nil {
			return ret, err
		}
		ret = append(ret, a)
	}
	return ret, nil
}
