package jsoninfo

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

// Config 구조체 선언
type JSONConfig struct {
	data map[string]interface{}
}

// NewConfig 함수 선언
func NewJSONConfig(fn string) (JSONConfigService, error) {

	// 현재 경로 출력
	pwd, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	// config 파일 불러오기
	data, err := os.Open(fmt.Sprintf("%s/%s", pwd, fn))
	if err != nil {
		return nil, err
	}

	// config 정보 받아오기
	byteData, err := io.ReadAll(data)
	if err != nil {
		return nil, err
	}

	var ji map[string]interface{}

	// 웹 서비스 정보에 담기
	json.Unmarshal(byteData, &ji)

	return &JSONConfig{data: ji}, nil

}
