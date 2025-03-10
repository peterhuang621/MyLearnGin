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
	// err := rdb.Set(ctx, "goredistestkey", "goredistestvalue", 10*time.Second).Err()
	// if err != nil {
	// 	panic(err)
	// }
	// val, err := rdb.Get(ctx, "goredistestkey").Result()
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println("goredistestkey:", val)

	// result, err := rdb.Do(ctx, "get", "goredistestkey").Result()
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println("do goredistextkey:", result.(string))

	// string
	// oldVal, err := rdb.GetSet(ctx, "goredistestkey", "new value").Result()
	// fmt.Println("key", oldVal)

	// err := rdb.SetNX(ctx, "key3", "value3", 0).Err()
	// vals, err := rdb.MGet(ctx, "key1", "key2", "key3").Result()
	// err := rdb.MSet(ctx, "key1", "v1", "key2", "v2", "key3", "v3").Err()
	// val, err := rdb.IncrBy(ctx, "k", 2).Result()
	// val, err := rdb.Expire(ctx, "key1", 5*time.Second).Result()

	// hashtable
	// err := rdb.HSet(ctx, "user_1", "username1", "huang").Err()
	// val, err := rdb.HGet(ctx, "user_1", "username1").Result()
	// val, err := rdb.HKeys(ctx, "user_1").Result()
	// val, err := rdb.HMGet(ctx, "user_1", "username", "username1").Result()
	// err := rdb.HMSet(ctx, "user_1", "username", "peter", "username1", "huanghuang").Err()
	// val, err := rdb.HExists(ctx, "user_1", "none").Result()

	// list
	// rdb.LPush(ctx, "lk", "data1", "data2")
	// rdb.LPushX(ctx, "lk", "data1", "data2")
	// val, err := rdb.LLen(ctx, "lk").Result()
	// val, err := rdb.LRange(ctx, "lk", 0, -1).Result()
	// err := rdb.LRem(ctx, "lk", 1, 2).Err()
	// val, err := rdb.LIndex(ctx, "lk", 2).Result()
	// err := rdb.LInsert(ctx, "lk", "before", "1", "2").Err()

	// set
	// err := rdb.SAdd(ctx, "ss", "aa", "bb", "aa").Err()
	// val, err := rdb.SMembers(ctx, "ss").Result()
	// err := rdb.SRem(ctx, "ss", "bb").Err()

	// sorted set
	// err := rdb.ZAdd(ctx, "zs", redis.Z{Score: 1.0, Member: "wangwu"}).Err()
	// val, err := rdb.ZCount(ctx, "zs", "1.3", "10.0").Result()
	// val, err := rdb.ZRange(ctx, "zs", 0.0, -1.0).Result()
	// val, err := rdb.ZRevRange(ctx, "zs", 0.0, -1.0).Result()
	// val, err := rdb.ZRangeByScoreWithScores(ctx, "zs", &redis.ZRangeBy{
	// 	Min:    "2.0",
	// 	Max:    "10.0",
	// 	Offset: 0,
	// 	Count:  3,
	// }).Result()
	// err := rdb.ZRemRangeByRank(ctx, "zs", 0, 1).Err()
	// val, err := rdb.ZRank(ctx, "zs", "huang").Result()

	// sub and con
	// sub := rdb.Subscribe(ctx, "channel1")
	// sub := rdb.PSubscribe(ctx, "channel*")
	// for ch := range sub.Channel() {
	// 	fmt.Println(ch.Channel)
	// 	fmt.Println(ch.Payload)
	// }
	// for {
	// 	message, err := sub.ReceiveMessage(ctx)
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// 	fmt.Println(message.Channel)
	// 	fmt.Println(message.Payload)
	// }
	// val, err := rdb.PubSubNumSub(ctx, "channel1").Result()

	// Multi Event

	// pipe := rdb.TxPipeline()
	// incr := pipe.Incr(ctx, "tx_pipeline_counter")
	// pipe.Expire(ctx, "tx_pipeline_counter", time.Hour)
	// _, err := pipe.Exec(ctx)
	// fmt.Println(incr.Val(), err)

	fn := func(tx *redis.Tx) error {
		v, err := tx.Get(ctx, "key").Int()
		if err != nil && err != redis.Nil {
			return err
		}
		v++

		_, err = tx.Pipelined(ctx, func(pipe redis.Pipeliner) error {
			pipe.Set(ctx, "key", v, 0)
			return nil
		})
		return err
	}

	for i := 0; i < 3; i++ {
		err := rdb.Watch(ctx, fn, "key")
		if err != nil {
			break
		}
		if err == redis.TxFailedErr {
			continue
		}
	}

	// fmt.Println(val)
}
