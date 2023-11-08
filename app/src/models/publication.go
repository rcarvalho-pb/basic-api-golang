package models

import (
	"database/sql"
)

type Publication struct {
	ID        int32 `json:"ID,omitempty"`
	Title     string `json:"Title,omitempty"`
	Content   string `json:"Content,omitempty"`
	AuthorID  uint64 `json:"AuthorID,omitempty"`
	Nick      string `json:"Nick,omitempty"`
	Likes     sql.NullInt32 `json:"Likes,omitempty"`
	CreatedAt sql.NullTime `json:"CreatedAt,omitempty"`
}
