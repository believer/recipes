package data

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

var (
	DB *sqlx.DB
)

func InitDB() error {
	db := sqlx.MustConnect("sqlite3", "./data/database.db")

	DB = db

	return nil
}
