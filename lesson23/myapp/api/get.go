package api

import (
	"github.com/labstack/echo/v4"
	"myapp/redis"
	"net/http"
)

func Get(c echo.Context) error {
	key := c.Param("key")
	response := redis.ClientGet(key)
	return c.String(http.StatusOK, response)
}