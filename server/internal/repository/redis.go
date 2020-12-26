package repository

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/pkg/errors"
	"github.com/wljgithub/mall-project/pkg/conf"
	"time"
)

const Nil = redis.Nil

func NewRedis() (*redis.Client, func(), error) {
	client, err := newRedis(conf.Conf.Redis)
	cf := func() {
		client.Close()
	}
	return client, cf, err
}

func newRedis(config conf.RedisConfig) (*redis.Client, error) {

	client := redis.NewClient(&redis.Options{
		Addr:         config.Addr,
		Password:     config.Password,
		DB:           config.Db,
		PoolSize:     config.PoolSize,
		DialTimeout:  time.Duration(config.DialTimeout) * time.Second,
		ReadTimeout:  time.Duration(config.ReadTimeout) * time.Second,
		WriteTimeout: time.Duration(config.WriteTimeout) * time.Second,
		PoolTimeout:  time.Duration(config.PoolTimeOut) * time.Second,
	})

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := client.Ping(ctx).Result()
	if err != nil {
		return nil, errors.Wrapf(err, "failed to ping redis")
	}
	return client, err
}
