package api

import (
	"github.com/labstack/echo/v4"
	"myapp/config"
	"myapp/db"
	"net/http"
)

func Post(c echo.Context) error{
	user := new(config.User)
	// body 에 들어온 값을 user에 Binding 한다.
	if err := c.Bind(&user); err != nil {
		return err
	}
	db.ClientSetRank(user)
	user.Rank = int(db.ClientGetRank(user.Username))
	return c.JSON(http.StatusOK, user)

}

