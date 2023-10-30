package models

import "database/sql"

type User struct {
	ID        int32
	Name      string
	Nick      string
	Email     string
	Password  string
	CreatedAt sql.NullTime
}
