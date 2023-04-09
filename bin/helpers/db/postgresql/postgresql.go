package postgresql

import (
	"database/sql"
	// "fmt"

	"github.com/cunkz/go-product/bin/config"
	_ "github.com/lib/pq"
)

var pgClient *sql.DB

func InitConnection() {
	psqlInfo := config.GetConfig().DsnPostgreSQL()
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	db.SetMaxOpenConns(5)
	db.SetMaxIdleConns(5)

	pgClient = db
}

func GetDB() *sql.DB {
	return pgClient
}
