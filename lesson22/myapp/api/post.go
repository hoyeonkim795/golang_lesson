package api

import (
	"github.com/labstack/echo/v4"
	"myapp/redis"
	"net/http"
)

func Post(c echo.Context) error {
	key := c.Param("key")
	value := c.Param("value")
	redis.ExampleClient(key, value)
	return c.String(http.StatusOK, "key:" + key + ", value:" + value)
}
