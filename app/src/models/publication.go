package models

import "time"

type Time struct {
	Time time.Time
	Valid bool
}

type Publication struct {
	ID int32 
	Title string 
	Content string 
	AuthorID uint64 
	Nick string 
	Likes int32 
	CreatedAt time.Time
}