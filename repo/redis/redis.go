package redisRepo

import (
	"context"
	"errors"
	"fmt"
	"github.com/redis/go-redis/v9"
	"time"
)

type Connection struct {
	client *redis.Client
}

func NewConnection() *Connection {
	return &Connection{
		// create redis connection
		client: redis.NewClient(&redis.Options{
			Addr: "localhost:6379",
		}),
	}
}

func (c *Connection) SetValue(ctx context.Context, key string, value interface{}, expired time.Duration) error {
	op := c.client.Set(ctx, key, value, expired)
	if err := op.Err(); err != nil {
		return errors.New(fmt.Sprintf("unable to SET data. error: %v", err))
	}
	return nil
}

func (c *Connection) GetValue(ctx context.Context, key string) (interface{}, error) {
	op := c.client.Get(ctx, key)
	if err := op.Err(); err != nil {
		return nil, errors.New(fmt.Sprintf("unable to GET data. error: %v", err))
	}
	res, err := op.Result()
	if err != nil {
		return nil, errors.New(fmt.Sprintf("unable to GET data. error: %v", err))
	}
	return res, nil
}
