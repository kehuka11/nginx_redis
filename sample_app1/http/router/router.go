package router

import (
	"sample_app1/http/handler"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

func NewRouter() *echo.Echo {
	e := echo.New()

	if l, ok := e.Logger.(*log.Logger); ok {
		l.SetLevel(log.INFO)
	}

	e.GET("/sample", handler.SampleApiHandler)

	return e
}
