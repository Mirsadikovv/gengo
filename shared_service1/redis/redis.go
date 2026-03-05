package redis

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

type Client redis.Client

var (
	instance *Client
)

func Instance(conf ...Config) *Client {

	if instance != nil {
		return instance
	}

	var config Config = defaultConfig
	{
		if len(conf) > 0 {
			config = conf[0]
		}
	}

	client, err := connect(config)
	{
		if err != nil {
			panic(err)
		}
	}

	instance = (*Client)(client)

	return instance
}

func connect(config Config) (*redis.Client, error) {

	rdb := redis.NewClient(
		&redis.Options{
			Addr: config.Addr,
		},
	)

	if _, err := rdb.Ping(context.Background()).Result(); err != nil {
		return nil, err
	}

	return rdb, nil
}

func (s *Client) SetWithTTL(key string, value any, timeOut time.Duration) error {

	status := s.Redis().Set(context.Background(), key, value, timeOut)
	{
		if err := status.Err(); err != nil {
			return err
		}
	}

	return nil
}

func (s *Client) Set(key string, value any) error {
	return s.SetWithTTL(key, value, 0)
}

func (s *Client) Get(key string) (any, error) {

	val, err := s.Redis().Get(context.Background(), key).Result()
	{
		if err != nil {
			return nil, err
		}
	}

	return val, nil
}

func (s *Client) Del(key string) error {
	return s.Redis().Del(context.Background(), key).Err()
}

func (s *Client) Redis() *redis.Client {
	return (*redis.Client)(s)
}
