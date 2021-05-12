package main

import (
	"fmt"
	"net/http"

	"context"

	"github.com/go-redis/redis/v8"


	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.POST("/save", save)

	e.Logger.Fatal(e.Start(":1323"))
}

func save(c echo.Context) error {
	key := c.FormValue("key")
	value := c.FormValue("value")
	ExampleClient(key, value)
	return c.String(http.StatusOK, "key:" + key + ", value:" + value)
}

var ctx = context.Background()

func ExampleClient(key string, value string) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	err := rdb.Set(ctx, key, value, 0).Err()
	if err != nil {
		panic(err)
	}

	val, err := rdb.Get(ctx, key).Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(key, val)

}