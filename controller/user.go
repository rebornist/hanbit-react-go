package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (r *Controller) UserController(c echo.Context) error {

	users, err := r.db.GetAllUsers()
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
