package models

import (
	"database/sql"
	"time"
)

type Time struct {
	Time  time.Time
	Valid bool
}

type Like struct {
	Int32 int32
	Valid bool
}

type Publication struct {
	ID        int32
	Title     string
	Content   string
	AuthorID  uint64
	Nick      string
	Likes     sql.NullInt32
	CreatedAt sql.NullTime
}
