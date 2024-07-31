package model

import "database/sql"

type Recipe struct {
	ID          int            `db:"id"`
	Description sql.NullString `db:"description"`
	Url         sql.NullString `db:"url"`
	Name        string         `db:"name"`
	Course      string         `db:"course"`
	Time        int            `db:"time"`
	Difficulty  string         `db:"difficulty"`
	Calories    sql.NullString `db:"calories"`
}

func (r Recipe) CourseType() string {
	switch r.Course {
	case "main":
		return "Huvudrätt"
	case "starter":
		return "Förrätt"
	default:
		return ""
	}
}

func (r Recipe) DifficultyText() string {
	switch r.Difficulty {
	case "easy":
		return "Lätt"
	case "medium":
		return "Mellan"
	case "hard":
		return "Svår"
	default:
		return ""
	}
}

type Ingredient struct {
	Name        string         `db:"ingredient"`
	Amount      sql.NullString `db:"amount"`
	ServingSize sql.NullString `db:"serving_size"`
}
