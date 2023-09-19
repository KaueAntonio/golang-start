package db

import (
	"context"
	"database/sql"
	"golang-start/App/App/Config"

	_ "github.com/go-sql-driver/mysql"
)

type Database struct {
	SqlDb *sql.DB
}

var dbContext = context.Background()

func OpenConnection() (*sql.DB, error) {
	cfg := Config.GetDbConnection()

	conn, err := sql.Open("mysql", cfg.FormatDSN())

	if err != nil {
		panic(err)
	}

	err = conn.Ping()

	return conn, err
}
