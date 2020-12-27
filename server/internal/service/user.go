package service

import (
	"context"
	"github.com/pkg/errors"
	"github.com/wljgithub/mall-project/internal/dto"
	"github.com/wljgithub/mall-project/internal/mapper"
	"github.com/wljgithub/mall-project/internal/model"
	"github.com/wljgithub/mall-project/pkg/conf"
	"github.com/wljgithub/mall-project/pkg/errno"
	"github.com/wljgithub/mall-project/pkg/token"
	"strconv"
	"time"
)

func (this *Service) Login(req dto.LoginReq) (*dto.LoginToken, error) {
	// 获取用户信息
	user, err := this.Repo.GetByName(context.Background(), req.LoginName)
	if err != nil {
		return nil, err
	}
	// 检查密码
	if !user.CheckPassword(req.PasswordMd5) {
		return nil, errors.Wrapf(errno.ErrIncorrectPassword, "%v password incorrect", user.LoginName)
	}

	// 签发 jwt token
	tokenStr, err := token.Sign(token.Context{UserID: uint64(user.UserId), Username: user.LoginName}, conf.Conf.App.JwtSecret)
	if err != nil {
		return nil, errors.Wrapf(err, "jwt token sign err")
	}

	// 将jwt token放入redis
	jwtExpire := time.Duration(conf.Conf.App.JwtExpire) * time.Hour
	if err := this.Repo.SetToken(context.Background(), strconv.Itoa(user.UserId), tokenStr, jwtExpire); err != nil {
		return nil, err
	}

	return &dto.LoginToken{Token: tokenStr}, nil

}

func (this *Service) Register(req dto.RegisterReq) error {
	return this.Repo.CreateUser(context.Background(), mapper.RegisterReqToUserModel(req))

}
func (this *Service) GetUserInfo(uid int) (*dto.GetUserInfoResp, error) {
	user, err := this.Repo.GetByUid(context.Background(), uid)
	if err != nil {
		return nil, err
	}
	return mapper.UserModelToUserInfo(user), nil

}

func (this *Service) UpdateUserInfo(uid int, req dto.UpdateUserInfoReq) error {
	return this.Repo.UpdateUser(&model.User{
		UserId:        uid,
		NickName:      req.NickName,
		IntroduceSign: req.IntroduceSign,
		PasswordMd5:   req.PasswordMd5,
	})
}
func (this *Service) Logout(uid int) error {
	return this.Repo.DeleteToken(context.Background(), uid)
}
