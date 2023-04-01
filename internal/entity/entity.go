package entity

import "database/sql"

type Task struct {
	Name   string
	Number int
	Text   string
}

type User struct {
	Name        string
	Level       int
	LastComment sql.NullString
}
