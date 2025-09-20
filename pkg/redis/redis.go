package redis

import (
	"context"
	"time"

	"online-library/pkg/config"

	"github.com/redis/go-redis/v9"
)

// Client struct for redis client
type Client struct {
	rdb *redis.Client
}

// new client for redis client
func NewClient(cfg config.Config) *Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     cfg.Redis.Addr,
		Password: cfg.Redis.Password,
		DB:       cfg.Redis.DB,
	})
	return &Client{rdb: rdb}
}

// set value in redis
func (c *Client) Set(ctx context.Context, key, value string, ttl time.Duration) error {
	return c.rdb.Set(ctx, key, value, ttl).Err()
}

// get value from redis
func (c *Client) Get(ctx context.Context, key string) (string, error) {
	return c.rdb.Get(ctx, key).Result()
}
