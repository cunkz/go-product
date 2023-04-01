package postgresql

import (
	"database/sql"
	// "fmt"

	"github.com/cunkz/go-lab-product/bin/config"
	// _ "github.com/lib/pq"
)

var pgClient *sql.DB

func InitConnection() {
	psqlInfo := config.GetConfig().DsnPostgreSQL()
	pgClient, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	pgClient.SetMaxOpenConns(5)
	pgClient.SetMaxIdleConns(5)

	defer pgClient.Close()
}

func GetDB() *sql.DB {
	return pgClient
}
