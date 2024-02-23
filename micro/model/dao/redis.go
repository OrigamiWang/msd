package dao

import (
	"context"
	"fmt"
	"time"

	"github.com/OrigamiWang/msd/micro/confparser"
	"github.com/redis/go-redis/v9"
)

type RedisClient struct {
	Client *redis.Client
}

func InitRedis(db *confparser.Database) *redis.Client {
	addr := fmt.Sprintf("%s:%d", db.Host, db.Port)
	return redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: db.Password,
		DB:       0,
		Username: db.User,
	})

}
func (rc *RedisClient) Get(key string) (interface{}, error) {
	return rc.Client.Get(context.Background(), key).Result()
}

func (rc *RedisClient) Set(key string, value any, expireTime time.Duration) error {
	return rc.Client.Set(context.Background(), key, value, expireTime).Err()
}

func (rc *RedisClient) Del(key string) error {
	return rc.Client.Del(context.Background(), key).Err()
}

func (rc *RedisClient) RPush(key string, value any) error {
	return rc.Client.RPush(context.Background(), key, value).Err()
}

func (rc *RedisClient) LRange(key string, start, end int64) ([]string, error) {
	return rc.Client.LRange(context.Background(), key, start, end).Result()
}

// RangeAll is a shortcut of LRange with start = 0, end = -1
func (rc *RedisClient) RangeAll(key string) ([]string, error) {
	return rc.Client.LRange(context.Background(), key, 0, -1).Result()
}

func (rc *RedisClient) Ttl(key string) (time.Duration, error) {
	return rc.Client.TTL(context.Background(), key).Result()
}

func (rc *RedisClient) Scan(match string, cnt int64) (keys []string, cursor uint64, err error) {
	return rc.Client.Scan(context.Background(), 0, match, cnt).Result()
}
