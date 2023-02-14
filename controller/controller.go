package controller

import (
	"database/sql"
	"hanbit-react/network"
	"hanbit-react/util/dblayer"
)

// HandlerService is a struct that implements ControllerInterface
type Controller struct {
	db  *sql.DB
	res network.Response
}

// NewHandler is a function that returns ControllerInterface
func NewController() (ControllerInterface, error) {
	// 기본 응답 객체 생성
	var res network.Response

	// Mysql 불러오기
	mysql, err := dblayer.NewDBConnection("mysql")
	if err != nil {
		return nil, err
	}

	db, err := mysql.MysqlConnection()
	if err != nil {
		return nil, err
	}

	return &Controller{db: db, res: res}, nil
}
