package postgresql

import (
	// "database/sql"
	"github.com/jmoiron/sqlx"
	// "fmt"

	"github.com/cunkz/go-product/bin/config"
	_ "github.com/lib/pq"
)

var pgClient *sqlx.DB

func InitConnection() {
	psqlInfo := config.GetConfig().DsnPostgreSQL()
	db, err := sqlx.Connect("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	db.SetMaxOpenConns(5)
	db.SetMaxIdleConns(5)

	pgClient = db
}

func GetDB() *sqlx.DB {
	return pgClient
}
