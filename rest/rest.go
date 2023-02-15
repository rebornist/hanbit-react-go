// Create Rest API with Echo Framework

package rest

import (
	"hanbit-react/view"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sirupsen/logrus"
)

func RunAPI(addr string) error {
	v, err := view.NewViewService()
	if err != nil {
		return err
	}
	return RunAPIWithHandler(addr, v)
}

func RunAPIWithHandler(addr string, v view.ViewSerivce) error {
	// Get Echo instance
	e := echo.New()

	// Logger
	log := logrus.New()

	// Middleware
	e.Use(middleware.RequestID())
	e.Use(middleware.Secure())
	e.Use(middleware.Recover())
	e.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogURI:    true,
		LogStatus: true,
		LogValuesFunc: func(c echo.Context, values middleware.RequestLoggerValues) error {
			log.WithFields(logrus.Fields{
				"URI":       values.URI,
				"method":    c.Request().Method,
				"connectIP": c.RealIP(),
				"status":    values.Status,
				"latency":   values.Latency,
				"agent":     c.Request().UserAgent(),
			}).Info("request")
			return nil
		},
	}))

	// Routes
	api := e.Group("/api")
	api.GET("/index", v.IndexView)
	api.GET("/users", v.LoginView)

	return e.Start(addr)
}
