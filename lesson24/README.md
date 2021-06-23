# Redis 랭킹 처리

## Requirments

- Redis 연동하기 2와 비슷한 패턴 활용으로 신규 기능을 추가한다.
- 신규 기능은 랭킹 처리 기능이다.
  - Redis의 Sorted Set을 활용하여 랭킹처리를 위한 get/set RESTful API를 구현한다.
  - get할 때, 랭킹의 범위를 함께 가져올 수 있도록 해야한다.
    - ex) GET rank/range/50   =>.  1 ~ 50 

## Main

```go
func main() {
	e := echo.New()
	e.GET("/ranklist/:start/:stop", api.GetRank)
	e.GET("/rank/:username", api.Get)
	e.POST("/", api.Post)

	e.Logger.Fatal(e.Start(":80"))
}
```


## POST

```go
func Post(c echo.Context) error{
	user := new(config.User)
	if err := c.Bind(&user); err != nil {
		return err
	}
	db.ClientSetRank(user)
	user.Rank = int(db.ClientGetRank(user.Username))
	return c.JSON(http.StatusOK, user)
}
```

- `POST` method '/' path로 body로 리퀘스트가 올 경우 해당 body 를 바인딩하여 user 인스턴스에 할당한다.
- 이후 Redis에서 `ClientSetRank` 메소드를 통해 데이터가 저장된다.
- 초기 저장시에 rank는 함께 body에 담아 요청하지 않고 radis에서 적재된 데이터들에서 rank가 정해지므로 body 를 redis에 적재한 뒤, redis내에서 rank를 확인하여 응답 메세지에 담아 보낸다.

## GET

### User의 Rank 확인

```go
func Get(c echo.Context) error{
	username := c.Param("username")
	rank := db.ClientGetRank(username)
	response := strconv.FormatInt(rank, 10)
	return c.String(http.StatusOK, response)
}
```

- /username path로 GET 요청이 올경우 path의 username을 받아 `CleintGetRank` 메소드의 파라미터 값으로 보낸다.
- `ClientGetRank` 를 통해서 rank 값을 가져오는데, 응답메세지의 경우 string 값으로 출력되기 때문에 `int64`타입의 rank 값을 string으로 변환한 뒤 응답메세지에 담아 보낸다.

### Ranklist 확인

```go
func GetRank(c echo.Context) error{
	start := c.Param("start")
	startNum, _ := strconv.ParseInt(start, 10, 64)
	stop := c.Param("stop")
	stopNum, _ := strconv.ParseInt(stop, 10, 64)
	rankarr,_ := db.ClientGetRankList(startNum, stopNum)
	return c.JSON(http.StatusOK, rankarr)
}
```

- 랭크를 확인하고 싶은 시작 index와 끝 index를 path로 받고, 해당 string 값을 `int64`타입으로 변환한다.
- 이후 `ClientGetRankList` 메소드로 해당 값들을 파라미터로 넘겨준다.
- 이후 받은 Range에 해당하는 user list 값을 Json 형태로 응답메시지로 보내준다.



# Client

## Set Post messagee

```go
func ClientSetRank(user *config.User) {

	client, err := NewDatabase()
	member := &redis.Z{
		Score: float64(user.Points),
		Member: user.Username,
	}
	if err != nil {
		fmt.Println(err)
	}
	client.ZAdd(ctx, "ranking", member)
}
```

- redis의 sorted set을 활용한다.
- 이후 POST 요청을 통해 들어온 user 인스턴스를 `ZAdd` 를 통해 저장한다.

## Get Rank

```go
func ClientGetRank(username string) int64 {
	client, err := NewDatabase()
	if err != nil {
		fmt.Println(err)
	}
	rank,_ := client.ZRank(ctx, "ranking", username).Result()
	return rank
}

```

- username에 해당하는 Rank를 응답으로 보낸다.
- redis의 `ZRank` 를 활용한다.

## Get Rank List

```go
func ClientGetRankList(start int64, stop int64) (*Leaderboard, error){
	client, err := NewDatabase()
	if err != nil {
		fmt.Println(err)
	}
	scores := client.ZRangeWithScores(ctx, "ranking", start, stop)
	if scores == nil {
		return nil, error(nil)
	}
	count := len(scores.Val())
	users := make([]*config.User, count)
	for idx, member := range scores.Val() {
		users[idx] = &config.User{
			Username: member.Member.(string),
			Points: int(member.Score),
			Rank: idx,
		}
	}
	leaderboard := &Leaderboard{
		Count: count,
		Users: users,
	}
	return leaderboard, nil
}
```

- path로 사용자가 확인하고 싶은 rank 순위 시작과 끝점을 파라미터로 받아 오름차순으로 해당하는 user 정보를 응답한다.
- `ZRangeWithScores` 를 활용해 해당 range에 해당하는 user 정보를 받는다.
- 그리고 `make` 내장함수를 통해 슬라이스로 range에 해당하는 user의 정보들을 순차적으로 담아 응답한다.



# Test

1. 서버 실행

```bash
$ myapp git:(master) ✗ make run
run server!

   ____    __
  / __/___/ /  ___
 / _// __/ _ \/ _ \
/___/\__/_//_/\___/ v4.3.0
High performance, minimalist Go web framework
https://echo.labstack.com
____________________________________O/_______
                                    O\
⇨ http server started on [::]:80

```

2. POST 요청

   ```bash
   $  myapp git:(master) ✗ http POST localhost:80/ <<< '{"username":"james kim", "points":13}'
   HTTP/1.1 200 OK
   Content-Length: 46
   Content-Type: application/json; charset=UTF-8
   Date: Tue, 22 Jun 2021 09:39:18 GMT
   
   {
       "points": 13,
       "rank": 3,
       "username": "james kim"
   }
   
   $ myapp git:(master) ✗ http POST localhost:80/ <<< '{"username":"elliott kim", "points":9}' 
   HTTP/1.1 200 OK
   Content-Length: 47
   Content-Type: application/json; charset=UTF-8
   Date: Tue, 22 Jun 2021 09:38:53 GMT
   
   {
       "points": 9,
       "rank": 1,
       "username": "elliott kim"
   }
   
   $  http POST localhost:80/ <<< '{"username":"elliott", "points":12}'
   HTTP/1.1 200 OK
   Content-Length: 44
   Content-Type: application/json; charset=UTF-8
   Date: Tue, 22 Jun 2021 09:38:38 GMT
   
   {
       "points": 12,
       "rank": 1,
       "username": "elliott"
   }
   
   $  myapp git:(master) ✗ http POST localhost:80/ <<< '{"username":"james", "points":15}'
   HTTP/1.1 200 OK
   Content-Length: 42
   Content-Type: application/json; charset=UTF-8
   Date: Tue, 22 Jun 2021 09:34:57 GMT
   
   {
       "points": 15,
       "rank": 0,
       "username": "james"
   }
   
   $  myapp git:(master) ✗ http POST localhost:80/ <<< '{"username":"bella", "points":3}' 
   HTTP/1.1 200 OK
   Content-Length: 41
   Content-Type: application/json; charset=UTF-8
   Date: Tue, 22 Jun 2021 09:35:06 GMT
   
   {
       "points": 3,
       "rank": 0,
       "username": "bella"
   }
   ```

   

3. Rank 확인

   ```bash
   $  myapp git:(master) ✗ http GET localhost:80/rank/bella                                   
   HTTP/1.1 200 OK
   Content-Length: 1
   Content-Type: text/plain; charset=UTF-8
   Date: Tue, 22 Jun 2021 10:07:32 GMT
   
   0
   
   ```

   

4. Rank list 확인

   ```bash
   $ myapp git:(master) ✗ http GET localhost:80/ranklist/0/4                                  
   HTTP/1.1 200 OK
   Content-Length: 196
   Content-Type: application/json; charset=UTF-8
   Date: Tue, 22 Jun 2021 09:38:58 GMT
   
   {
       "Users": [
           {
               "points": 3,
               "rank": 0,
               "username": "bella"
           },
           {
               "points": 9,
               "rank": 1,
               "username": "elliott kim"
           },
           {
               "points": 12,
               "rank": 2,
               "username": "elliott"
           },
           {
               "points": 15,
               "rank": 3,
               "username": "james"
           }
       ],
       "count": 4
   }
   
   
   ```

   