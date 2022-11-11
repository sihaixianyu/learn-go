package types

import (
	"context"

	"github.com/go-redis/redis/v9"
)

var ctx = context.Background()
var client *redis.Client

func init() {
	client = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
}

func StringExample() {
	err := client.Set(ctx, "name", "tingxuan", 0).Err()
	if err != nil {
		panic(err)
	}
}
