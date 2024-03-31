package sqlite

import (
	"database/sql"
	"matthewhope/atm-system-go/models"
)

func GetUsers(db *sql.DB) ([]models.User, error) {
	var ret []models.User
	query := `SELECT * FROM "USERS";`
	rows, err := db.Query(query)
	if err != nil {
		return ret, nil
	}
	for rows.Next() {
		var u models.User
		err = rows.Scan(&u.UserID, &u.CreatedAt, &u.EmailAddress, &u.EncryptedPassword, &u.FirstName, &u.LastName, &u.UpdatedAt, &u.Username)
		if err != nil {
			return ret, err
		}
		ret = append(ret, u)
	}
	return ret, nil
}
