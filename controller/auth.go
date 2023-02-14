package controller

import (
	"hanbit-react/auth"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (r *Controller) AuthController(c echo.Context) error {

	// 인증 서비스 생성
	a, err := auth.NewAuthService()
	if err != nil {
		r.res.Code = http.StatusInternalServerError
		r.res.Msg = "Internal Server Error"
		r.res.Data = nil
		return c.JSON(http.StatusInternalServerError, r.res)
	}

	// 유저 목록 가져오기
	users, err := a.GetAllUsers()
	if err != nil {
		r.res.Code = http.StatusInternalServerError
		r.res.Msg = "Internal Server Error"
		r.res.Data = nil
		return c.JSON(http.StatusInternalServerError, r.res)
	}

	r.res.Code = http.StatusOK
	r.res.Msg = "Success"
	r.res.Data = map[string]interface{}{"title": users}

	return c.JSON(http.StatusOK, r.res)

}
