# Json POST & GET

## Redis Installation

  - https://github.com/go-redis/redis

  - Redis port : 6379 

  - ```bash
    $ go mod init github.com/my/repo
    $ go get github.com/go-redis/redis/v8
    ```

## Post

### API

```go
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
```

- two-depth를 가진 Name은 Marshal해 json 형태로 출력하도록 함
- User struct의 FullName, Age, Hobby 의 parameter를 `Client` 함수로 전달

### Redis

```go
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
```

- 넘겨 받은 파라미터 값을 `HMSet`으로 redis에 저장

## Get

### API

```go
func Get(c echo.Context) error {
	key := c.Param("key")
	response := redis.ClientGet(key)
	return c.String(http.StatusOK, response)
}
```

- ClientGet함수로 path로 받은 `key` 값을 넘겨줌

### Client

```go
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
```

- FullName의 경우 two depth를 갖기 때문에 Unmarshal을 통해 firstname과 lastname 값을 name에 할당
- 각 필드에 해당하는 값을 할당 후 Marshal 한 뒤 해당 값을 `return`



## Makefile

- MakeFile을 통한 Echo Framework 실행
  - ```$ make run```
- Redis Cli 연결
  - ```make show redis```

## Test

### httpie

1. Installation

   ```bash
   $ <tool> install httpie
   ```

2. Post Key, Value

   ```bash
   $ http POST locallhost:1323/<key>/<value>
   ```

3. Check Key, Value

   ```$ get <key>```

### Example

1. Echoframework 실행

   ```$ make run```

   ```bash
   run server!
   
      ____    __
     / __/___/ /  ___
    / _// __/ _ \/ _ \
   /___/\__/_//_/\___/ v4.3.0
   High performance, minimalist Go web framework
   https://echo.labstack.com
   ____________________________________O/_______
                                       O\
   ⇨ http server started on [::]:1323
   
   ```

   

2. httpie를 이용한 POST 요청 테스트

```bash
$ http POST localhost:1323/123 <<< '{"name" :{"first_name":"any","last_name":"kim"},"age":"1","hobby":"run"}'
res0: HTTP/1.1 200 OK
Content-Length: 72
Content-Type: application/json; charset=UTF-8
Date: Mon, 07 Jun 2021 09:32:36 GMT

{
    "age": "1",
    "hobby": "run",
    "name": {
        "first_name": "any",
        "last_name": "kim"
    }
}
```

​		

3.  httpie를 이용한 GET 요청 테스트

   ```bash
   $ http GET localhost:1323/123    
   res0:HTTP/1.1 200 OK
   Content-Length: 71
   Content-Type: text/plain; charset=UTF-8
   Date: Mon, 07 Jun 2021 09:52:36 GMT
   
   {
       "age": "1",
       "hobby": "run",
       "name": {
           "first_name": "any",
           "last_name": "kim"
       }
   }
   
   ```

   

make 명령어로 Redis Cli 열고 key, value 확인

```bash
$ make show redis
$ HMGET <key> <field> // key 의 field value를 얻고 싶을때
$ HGETALL <key> // key 의 모든 field와 value를 얻고 싶을때
```

```bash
$ HGET 123 "hobby"
res0: 1) "run"
$ HGET 123 "name"
res1: 1) "{\"first_name\":\"any\",\"last_name\":\"kim\"}"
$ HGETALL 123
res2: 1) "name"
2) "{\"first_name\":\"any\",\"last_name\":\"kim\"}"
3) "age"
4) "1"
5) "hobby"
6) "run"

```

