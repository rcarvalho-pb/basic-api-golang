package db

import (
	"api/db/database"
	"api/src/config"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func Conn() (*sql.DB, *database.Queries, error) {
	db, err := sql.Open("mysql", config.ConnString)
	if err != nil {
		return nil, nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, nil, err
	}
	
	queries := database.New(db)

	return db, queries, nil
}