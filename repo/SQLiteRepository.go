package repo

import (
	"database/sql"
	"log"
	"matthewhope/atm-system-go/sqlite"

	_ "github.com/mattn/go-sqlite3"
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

func (r *SQLiteRepository) UsersTableUp() error {
	return sqlite.UsersTableUp(r.identityDb)
}
