package view

import (
	"github.com/labstack/echo/v4"
)

type ViewSerivce interface {
	// Index Page
	IndexView(e echo.Context) error

	// Auth
	LoginView(e echo.Context) error
}
