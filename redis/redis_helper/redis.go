package redis_helper

import (
	"context"

	"github.com/go-redis/redis/v8"
)

type Redis interface {
	Read(target string) (content string, err error)
	Write(target string, content string) (err error)
}

type redisInstance struct {
	rdb *redis.Client
	ctx context.Context
}

func NewRedis(ctx context.Context) Redis {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	return &redisInstance{
		rdb,
		ctx,
	}
}

func (r *redisInstance) Read(target string) (content string, err error) {
	content, err = r.rdb.Get(r.ctx, target).Result()
	return content, err
}

func (r *redisInstance) Write(target string, content string) (err error) {
	_, err = r.rdb.Set(r.ctx, target, content, 0).Result()
	return err
}
