package dblayer

import "database/sql"

type DBService interface {
	MysqlConnection() (*sql.DB, error)
}
