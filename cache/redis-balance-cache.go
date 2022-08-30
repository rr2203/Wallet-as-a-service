package cache

import (
	"context"
	"github.com/go-redis/redis/v8"
	"time"
)

type redisCache struct {
	host string
	db int
	expires time.Duration
}

func NewRedisCache(host string, db int, expires time.Duration) BalanceCache {
	return &redisCache{
		host: host,
		db: db,
		expires: expires,
	}
}

func (cache *redisCache) getClient() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr: cache.host,
		Password: "",
		DB: cache.db,
	})
}

func (cache *redisCache) SET(key int, value float64) error{
	client := cache.getClient()
	result := client.Set(context.TODO(), string(rune(key)), value, 0)
	if result.Err() != nil	{
		return result.Err()
	}
	return nil
}

func (cache *redisCache) GET(key int) (string, error){
	client := cache.getClient()
	result := client.Get(context.TODO(), string(rune(key)))
	if result.Err() != nil	{
		return result.Val(), result.Err()
	}
	return result.Val(), nil
}