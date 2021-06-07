package api

import (
	"encoding/json"
	"github.com/labstack/echo/v4"
	"myapp/redis"
	"myapp/struct"
	"net/http"
)


func Save(c echo.Context) (err error) {
	key := c.Param("key")
	user := new(_struct.User)
	// body 에 들어온 값을 user에 Binding 한다.
	if err = c.Bind(user); err != nil {
		return
	}
	fullName, _ := json.Marshal(user.FullName)
	redis.Client(key, "name", string(fullName), "age", user.Age, "hobby", user.Hobby )
	return c.JSON(http.StatusOK, user)
}