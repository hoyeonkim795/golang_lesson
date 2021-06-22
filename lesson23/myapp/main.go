package main

import (
	"github.com/labstack/echo/v4"
	"myapp/api"
)

func main() {

	e := echo.New()

	e.GET("/:key", api.Get)
	e.POST("/:key", api.Save)

	e.Logger.Fatal(e.Start(""))

}