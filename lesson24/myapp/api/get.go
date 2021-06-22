package api

import (
	"github.com/labstack/echo/v4"
	"myapp/db"
	"net/http"
	"strconv"
)

func Get(c echo.Context) error{
	username := c.Param("username")
	rank := db.ClientGetRank(username)
	response := strconv.FormatInt(rank, 10)
	return c.String(http.StatusOK, response)
}


func GetRank(c echo.Context) error{
	start := c.Param("start")
	startNum, _ := strconv.ParseInt(start, 10, 64)
	stop := c.Param("stop")
	stopNum, _ := strconv.ParseInt(stop, 10, 64)
	rankarr,_ := db.ClientGetRankList(startNum, stopNum)
	return c.JSON(http.StatusOK, rankarr)
}