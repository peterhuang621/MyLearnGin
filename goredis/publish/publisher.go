package main

import (
	"context"

	"github.com/redis/go-redis/v9"
)

var rdb *redis.Client

func init() {
	rdb = redis.NewClient(&redis.Options{Addr: "localhost:6379", Password: "", DB: 0})
}

func main() {
	ctx := context.Background()
	rdb.Publish(ctx, "channel1", "msg")
}
