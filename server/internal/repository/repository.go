package repository

import (
	"github.com/go-redis/redis/v8"
	"github.com/google/wire"
	"gorm.io/gorm"
)

type Repository interface {
	Close()
}

var _ Repository = &Repo{}

var Provider = wire.NewSet(New, NewRedis, NewMysql)

type Repo struct {
	db    *gorm.DB
	redis *redis.Client
}

func New(redis *redis.Client, db *gorm.DB) (Repository, func(), error) {
	repository := &Repo{
		db:    db,
		redis: redis,
	}
	cf := func() {
		repository.Close()
	}
	return repository, cf, nil
}

func (this *Repo) Close() {

}
