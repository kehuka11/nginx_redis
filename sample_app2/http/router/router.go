package router

import (
	"sample_app2/http/handler"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

func NewRouter() *echo.Echo {
	e := echo.New()

	e.GET("/sample", handler.SampleApiHandler)

	if l, ok := e.Logger.(*log.Logger); ok {
		l.SetLevel(log.INFO)
	}

	return e
}
