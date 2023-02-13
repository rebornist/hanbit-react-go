package model

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

// Connection info for the DB
type DBLayer struct {
	*sql.DB
}

// Connect to the Mysql
func NewDBLayer(dsn string) (*DBLayer, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	return &DBLayer{db}, nil
}
