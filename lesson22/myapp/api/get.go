package api

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func Get(c echo.Context) error {
	return c.String(http.StatusOK, "hello world")
}