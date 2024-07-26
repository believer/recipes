package model

import "database/sql"

type Recipe struct {
	ID          int            `db:"id"`
	Description sql.NullString `db:"description"`
	Url         sql.NullString `db:"url"`
	Name        string         `db:"name"`
	Type        string         `db:"type"`
}

type Ingredient struct {
	Name   string `db:"name"`
	Amount string `db:"amount"`
}
