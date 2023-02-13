package controller

import (
	"github.com/labstack/echo/v4"
)

type ControllerInterface interface {
	// Index Page
	IndexController(e echo.Context) error

	// User Page
	UserController(e echo.Context) error
}
