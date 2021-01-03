package repository

import (
	"errors"
	"github.com/go-redis/redis/v8"
	"github.com/google/wire"
	"github.com/wljgithub/mall-project/internal/model"
	"gorm.io/gorm"
)

type Repository interface {
	UserRepo
	MallRepo
	AddressRepo
	CartRepo
	OrderRepo
	Close()
}

var tableToRegister = []interface{}{
	&model.Order{},
	&model.GoodsItem{},
}
var (
	ErrNotFound = errors.New("record not found")
)

var RedisClient *redis.Client

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
	RedisClient = redis
	cf := func() {
		repository.Close()
	}
	return repository, cf, nil
}

func (this *Repo) Close() {

}
