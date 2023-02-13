package controller

import (
	"encoding/json"
	"fmt"
	"hanbit-react/model"
	"hanbit-react/network"
	"hanbit-react/types"
	ji "hanbit-react/util/jsonInfo"
)

// HandlerService is a struct that implements ControllerInterface
type Controller struct {
	db  *model.DBLayer
	res network.Response
}

// NewHandler is a function that returns ControllerInterface
func NewController() (ControllerInterface, error) {
	// 데이터베이스 타입 초기화
	var db types.Database

	// config 핸들러 선언
	config, err := ji.NewJSONConfig("config_local.json")
	if err != nil {
		return nil, err
	}

	// database 접속 정보 불러오기
	di, err := config.GetJSONInfo("database")
	if err != nil {
		return nil, err
	}

	// 데이터베이스 타입에 링크
	err = json.Unmarshal(di, &db)
	if err != nil {
		return nil, err
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local&autocommit=false",
		db.User,
		db.Password,
		db.Host,
		db.Port,
		db.Name,
	)

	return NewHandlerWithParams(dsn)
}

func NewHandlerWithParams(dsn string) (ControllerInterface, error) {

	// 기본 응답 객체 생성
	var res network.Response

	// gorm 불러오기
	db, err := model.NewDBLayer(dsn)
	if err != nil {
		return nil, err
	}

	return &Controller{db: db, res: res}, nil
}
