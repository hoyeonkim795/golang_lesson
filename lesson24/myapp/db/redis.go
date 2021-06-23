package db

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"myapp/config"
)

type Database struct {
	Client *redis.Client
}

type Leaderboard struct {
	Count int `json:"count"`
	Users []*config.User
}

var ctx = context.Background()

func NewDatabase() (*redis.Client, error){
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	_, err := client.Ping(ctx).Result()

	return client, err

}

func ClientSetRank(user *config.User) {

	client, err := NewDatabase()
	member := &redis.Z{
		Score: float64(user.Points),
		Member: user.Username,
	}
	//pipe := client.TxPipeline()
	//pipe.ZAdd(ctx,"leader", member)
	//rank := pipe.ZRank(ctx, "leader", user.Username)
	//_, err = pipe.Exec(ctx)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//user.Rank = int(rank.Val())
	if err != nil {
		fmt.Println(err)
	}
	client.ZAdd(ctx, "ranking", member)
}

func ClientGetRank(username string) int64 {
	client, err := NewDatabase()
	if err != nil {
		fmt.Println(err)
	}
	//pipe := client.TxPipeline()
	//rank, _ := pipe.ZRank(ctx, "leader", username).Result()
	rank,_ := client.ZRank(ctx, "ranking", username).Result()
	return rank
}

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
			Rank: idx+int(start),
		}
	}
	leaderboard := &Leaderboard{
		Count: count,
		Users: users,
	}
	return leaderboard, nil
}