package distributed_lock

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
	"time"
)

type Client struct {
	client redis.Cmdable
}

func NewClient(c redis.Cmdable) *Client {
	return &Client{
		client: c,
	}
}

type Lock struct {
	key   string
	value string
}

func (c *Client) Lock(ctx context.Context, key string) (*Lock, error) {

	val := uuid.New().String()
	res, err := c.client.SetNX(ctx, key, val, time.Minute).Result()
	if err != nil {
		return nil, err
	}
	if !res {
		return nil, errors.New("解锁失败")
	}
	return &Lock{
		key:   key,
		value: val,
	}, nil
}
