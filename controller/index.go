package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (r *Controller) IndexController(c echo.Context) error {

	r.res.Code = http.StatusOK
	r.res.Msg = "Success"
	r.res.Data = map[string]interface{}{"title": "Index Page"}

	return c.JSON(http.StatusOK, r.res)

}
