package dblayer

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"hanbit-react/types"
	"hanbit-react/util/component"

	_ "github.com/go-sql-driver/mysql"
)

// DBLayer is a struct that implements DBLayerInterface
func (d *DB) MysqlConnection() (*sql.DB, error) {
	dsn, err := getMysqlConnStr()
	if err != nil {
		return nil, err
	}

	db, err := sql.Open(d.cate, dsn)
	if err != nil {
		return nil, err
	}
	return db, nil
}

// Mysql Connection String 가져오기
func getMysqlConnStr() (string, error) {
	// 데이터베이스 타입 초기화
	var db types.Database

	// config 핸들러 선언
	conf, err := component.NewJSONConfig("config_local.json")
	if err != nil {
		return "", err
	}

	// database 접속 정보 불러오기
	di, err := conf.GetServiceInfo("database")
	if err != nil {
		return "", err
	}

	// 데이터베이스 타입에 링크
	err = json.Unmarshal(di, &db)
	if err != nil {
		return "", err
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local&autocommit=false",
		db.User,
		db.Password,
		db.Host,
		db.Port,
		db.Name,
	)

	return dsn, nil
}
