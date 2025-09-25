package redis

import (
	"context"

	"github.com/redis/go-redis/v9"
)

type rdb struct {
	client *redis.Client
}

func newClient() *rdb {
	reddb:=redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
	return &rdb{client:reddb}
}

func (rdb *rdb ) SubscribeTORedis(ctx context.Context, channel string)*redis.PubSub{
	sub:=rdb.client.Subscribe(ctx,channel)
	return sub
}

func (rdb *rdb ) PublishToRedis(ctx context.Context, channel string, data []byte) *redis.IntCmd{
	pub:=rdb.client.Publish(ctx, channel, data)
	return pub
}

func (rdb *rdb) Set(ctx context.Context, key string, members ...interface{}) *redis.IntCmd{
	set:= rdb.client.SAdd(ctx, key, members)
	return set
}

func (rdb *rdb) UnSet(ctx context.Context, key string, members ...interface{}) *redis.IntCmd{
	unset:= rdb.client.SRem(ctx, key, members)
	return unset
}
