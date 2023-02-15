package view

import (
	"hanbit-react/network"
)

type ViewResponse struct {
	network.Response
}

// 뷰 서비스 생성
func NewViewService() (ViewSerivce, error) {
	// 기본 응답 객체 생성
	var res network.Response

	return &ViewResponse{res}, nil
}
