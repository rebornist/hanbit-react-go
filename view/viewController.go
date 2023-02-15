package view

import (
	"hanbit-react/auth"
	"net/http"

	"github.com/labstack/echo/v4"
)

// Index Page
func (res *ViewResponse) IndexView(c echo.Context) error {

	res.Code = http.StatusOK
	res.Msg = "Success"
	res.Data = map[string]interface{}{"title": "Index Page"}

	return c.JSON(http.StatusOK, res)

}

func (res *ViewResponse) LoginView(c echo.Context) error {

	// // 인증 서비스 생성
	// u, err := user.NewUserService()
	// if err != nil {
	// 	res.Code = http.StatusInternalServerError
	// 	res.Msg = "Internal Server Error"
	// 	res.Data = nil
	// 	return c.JSON(http.StatusInternalServerError, res)
	// }

	// // 유저 목록 가져오기
	// users, err := u.GetAllUsers()
	// if err != nil {
	// 	res.Code = http.StatusInternalServerError
	// 	res.Msg = "Internal Server Error"
	// 	res.Data = nil
	// 	return c.JSON(http.StatusInternalServerError, res)
	// }

	a, err := auth.NewAuthService()
	if err != nil {
		res.Code = http.StatusInternalServerError
		res.Msg = "Internal Server Error"
		res.Data = map[string]interface{}{"message": err.Error()}
		return c.JSON(http.StatusInternalServerError, res)
	}

	user, err := a.BasicAuthentication("test", "test")
	if err != nil {
		res.Code = http.StatusInternalServerError
		res.Msg = "Internal Server Error"
		res.Data = map[string]interface{}{"message": err.Error()}
		return c.JSON(http.StatusInternalServerError, res)
	}

	res.Code = http.StatusOK
	res.Msg = "Success"
	res.Data = map[string]interface{}{"title": user}

	return c.JSON(http.StatusOK, res)

}
