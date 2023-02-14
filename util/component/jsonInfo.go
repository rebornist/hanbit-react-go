package component

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

// 핸들러 전체 메서드 포함 인터페이스
type JSONConfigInterface interface {
	GetServiceInfo(name string) ([]byte, error)
}

// Config 구조체 선언
type JSONConfig struct {
	Data map[string]interface{}
}

// NewJSONConfig 함수 선언
func NewJSONConfig(fn string) (JSONConfigInterface, error) {

	// 현재 경로 출력
	pwd, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	// config 파일 불러오기
	conf, err := os.Open(fmt.Sprintf("%s/%s", pwd, fn))
	if err != nil {
		return nil, err
	}

	// config 정보 받아오기
	byteData, err := io.ReadAll(conf)
	if err != nil {
		return nil, err
	}

	var data map[string]interface{}

	json.Unmarshal(byteData, &data)

	return &JSONConfig{Data: data}, nil

}

func (c *JSONConfig) GetServiceInfo(name string) ([]byte, error) {
	return json.Marshal(c.Data[name])
}
