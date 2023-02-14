package auth

import (
	"database/sql"
	"hanbit-react/network"
	"hanbit-react/util/dblayer"
)

// Connection info for the DB
type Auth struct {
	db  *sql.DB
	res network.Response
}

// NewHandler is a function that returns ControllerInterface
func NewAuthService() (AuthSerivce, error) {

	// 기본 응답 객체 생성
	var res network.Response

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

	return &Auth{db: db, res: res}, nil
}
