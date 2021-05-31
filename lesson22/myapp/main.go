package main

import (
	"github.com/labstack/echo/v4"
	"myapp/api"
)

func main() {
	e := echo.New()

	e.GET("/", api.Get)
	e.POST("/:key/:value", api.Post)
	e.Logger.Fatal(e.Start(":1323"))
}