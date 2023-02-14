package controller

import (
	"github.com/labstack/echo/v4"
)

type ControllerInterface interface {
	// Index Page
	IndexController(e echo.Context) error

	// Auth
	AuthController(e echo.Context) error
}
