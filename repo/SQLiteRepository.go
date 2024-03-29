package repo

import (
	"database/sql"
	"log"
)

const (
	dbDriver       = "sqlite3"
	identityDbPath = "./sqlite/data/IDENTITY.db"
	businessDbPath = "./sqlite/data/BUSINESS.db"
)

type SQLiteRepository struct {
	identityDb, businessDb *sql.DB
}

func NewSQLiteRepository() *SQLiteRepository {
	var err error
	ret := new(SQLiteRepository)
	ret.identityDb, err = sql.Open(dbDriver, identityDbPath)
	if err != nil {
		log.Fatal(err.Error())
	}
	ret.businessDb, err = sql.Open(dbDriver, businessDbPath)
	if err != nil {
		log.Fatal(err.Error())
	}
	return ret
}
