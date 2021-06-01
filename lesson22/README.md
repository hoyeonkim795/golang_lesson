# Echo Framework Redis 연동

## Redis Installation
  - https://github.com/go-redis/redis 을 참고
  - 6379 port에 Redis 연결
  - ```
    go mod init github.com/my/repo
    go get github.com/go-redis/redis/v8
     ```
## Client
```bigquery
func Client(key string, value string) {
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
```
## API

- path 로 전달되는 파라미터를 획득 (https://lejewk.github.io/go-echo-request-param/)
- 넘겨준 파라미터 (key, value)를 작성한 `ExampleClient`로 넘김
- ```bigquery
    func Post(c echo.Context) error {

	    key := c.Param("key")
	    value := c.Param("value")
	    redis.Client(key, value)
	    return c.String(http.StatusOK, "key:" + key + ", value:" + value)
    }   

    ```

## Test
- httpie 를 사용 : https://httpie.io/
- e.g, key: bella, value: spoon 테스트
- ` http POST localhost:1323/bella/spoon`
- get [key] 를 통해 value 확인

## Makefile
- MakeFile을 통한 Redis Cli, Echo Framework 실행