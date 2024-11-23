package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func SampleApiHandler(c echo.Context) error {
	message := "アプリ2です"

	return c.String(http.StatusOK, message)
}
