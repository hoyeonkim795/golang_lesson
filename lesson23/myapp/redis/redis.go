package redis

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis/v8"
	_struct "myapp/struct"
)

var ctx = context.Background()


func Client(key string, field1 string, value1 string, field2 string, value2 string, field3 string, value3 string) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	err := rdb.HMSet(ctx, key, field1, value1, field2, value2, field3, value3).Err()
	if err != nil {
		panic(err)
	}
	fmt.Println(key)

	if err != nil {
		fmt.Println(err)
	}
}


func ClientGet(key string) string{


	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		Password: "",
		DB: 0,
	})
	fmt.Println(rdb)

	val , err := rdb.HMGet(ctx, key, "name", "age", "hobby").Result()
	if err != nil {
		panic(err)
	}
	name := new(_struct.Name)
	err2 := json.Unmarshal([]byte(val[0].(string)), &name)
	if err2 != nil {
		fmt.Println(err2)
	}

	dat, _ := json.Marshal(_struct.User{FullName: _struct.Name{FirstName: name.FirstName, LastName: name.LastName}, Age: val[1].(string), Hobby: val[2].(string)})

	return string(dat)
}
