package cache

import (
	"context"
	"os"
	"time"

	"github.com/go-redis/redis/v9"
)

var redisCtx context.Context

var rdb *redis.Client

type RedisRepo struct{}

func init() {
	redisCtx = context.Background()
	rdb = redis.NewClient(&redis.Options{
		Addr:     os.Getenv("CACHE_HOST") + ":" + os.Getenv("CACHE_PORT"),
		Password: "",
		DB:       0,
	})
}

func (r *RedisRepo) New() *RedisRepo {
	return &RedisRepo{}
}

// Set value to redis, Always KeepTTL
func (r *RedisRepo) Set(key string, value interface{}) error {
	return rdb.SetArgs(redisCtx, key, value, redis.SetArgs{
		KeepTTL: true,
	}).Err()
}

// Setex
func (r *RedisRepo) Setex(key string, value interface{}, ttl time.Duration) error {
	return rdb.SetEx(redisCtx, key, value, ttl).Err()
}

func (r *RedisRepo) Get(key string) (string, error) {
	return rdb.Get(redisCtx, key).Result()
}

func (r *RedisRepo) GetJson(key string, des interface{}) error {
	return rdb.Get(redisCtx, key).Scan(des)
}

func (r *RedisRepo) Del(key string) error {
	return rdb.Del(redisCtx, key).Err()
}

func (r *RedisRepo) Increase(key string) error {
	return rdb.Incr(redisCtx, key).Err()
}

func (r *RedisRepo) Decrease(key string) error {
	return rdb.Decr(redisCtx, key).Err()
}
