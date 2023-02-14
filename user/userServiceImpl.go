package user

import (
	"database/sql"
	"hanbit-react/util/dblayer"
)

// Connection info for the DB
type User struct {
	*sql.DB
}

// NewHandler is a function that returns ControllerInterface
func NewUserService() (UserSerivce, error) {

	// Mysql 불러오기
	mysql, err := dblayer.NewDBConnection("mysql")
	if err != nil {
		return nil, err
	}

	// Mysql Connection
	db, err := mysql.MysqlConnection()
	if err != nil {
		return nil, err
	}

	return &User{db}, nil
}
