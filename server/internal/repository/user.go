package repository

import (
	"context"
	"github.com/go-sql-driver/mysql"
	xerrors "github.com/pkg/errors"
	"github.com/wljgithub/mall-project/internal/model"
	"github.com/wljgithub/mall-project/pkg/errno"
	"strconv"
	"time"
)

type UserRepo interface {
	GetByUid(ctx context.Context, uid int) (*model.User, error)
	GetByName(ctx context.Context, name string) (*model.User, error)
	SetToken(ctx context.Context, uid string, tokenStr string, expire time.Duration) error
	CreateUser(ctx context.Context, user model.User) error
	UpdateUser(u *model.User) error
	DeleteToken(ctx context.Context, uid int) error
}

func (this *Repo) GetByUid(ctx context.Context, uid int) (*model.User, error) {
	user := &model.User{}
	err := this.db.Where(&model.User{UserId: uid}).First(user).Error
	return user, err
}

func (this *Repo) GetByName(ctx context.Context, name string) (*model.User, error) {
	user := &model.User{}
	err := this.db.Where(&model.User{LoginName: name}).First(user).Error
	if err != nil {
		return nil, xerrors.Wrapf(err, "failed to get user by name")
	}
	return user, err
}
func (this *Repo) SetToken(ctx context.Context, uid string, tokenStr string, expire time.Duration) error {
	err := this.redis.Set(ctx, uid, tokenStr, expire).Err()
	if err != nil {
		return xerrors.Wrapf(err, "failed to set jwt token in redis")
	}
	return nil
}
func (this *Repo) CreateUser(ctx context.Context, user model.User) error {
	err := this.db.Create(&user).Error
	if typed, ok := err.(*mysql.MySQLError); ok && typed.Number == 1062 {
		return xerrors.Wrapf(errno.ErrUserExist, "")
	}
	return err
}
func (this *Repo) UpdateUser(u *model.User) error {
	return this.db.Model(u).Where("user_id", u.UserId).Updates(*u).Error
}
func (this *Repo) DeleteToken(ctx context.Context, uid int) error {
	err := this.redis.Del(ctx, strconv.Itoa(uid)).Err()
	return xerrors.Wrapf(err, "failed to delete token in redis")
}
