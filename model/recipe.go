package model

import "database/sql"

type Recipe struct {
	ID          int            `db:"id"`
	Description sql.NullString `db:"description"`
	Url         sql.NullString `db:"url"`
	Name        string         `db:"name"`
	Course      string         `db:"course"`
}

func (r Recipe) CourseType() string {
	switch r.Course {
	case "main":
		return "Huvudrätt"
	default:
		return ""
	}
}

type Ingredient struct {
	Name        string         `db:"ingredient"`
	Amount      sql.NullString `db:"amount"`
	ServingSize sql.NullString `db:"serving_size"`
}
