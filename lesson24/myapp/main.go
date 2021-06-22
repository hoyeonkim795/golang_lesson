package main

import (
	"github.com/labstack/echo/v4"
	"myapp/api"
)

func main() {
	e := echo.New()
	e.GET("/ranklist/:start/:stop", api.GetRank)
	e.GET("/rank/:username", api.Get)
	e.POST("/", api.Post)

	e.Logger.Fatal(e.Start(":80"))

}