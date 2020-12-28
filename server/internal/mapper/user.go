package mapper

import (
	"github.com/wljgithub/mall-project/internal/dto"
	"github.com/wljgithub/mall-project/internal/model"
	"time"
)

func UserModelToUserInfo(user *model.User) *dto.GetUserInfoResp {
	return &dto.GetUserInfoResp{
		IntroduceSign: user.IntroduceSign,
		LoginName:     user.LoginName,
		NickName:      user.NickName,
	}
}

func RegisterReqToUserModel(req dto.RegisterReq) model.User {
	return model.User{
		LoginName:   req.LoginName,
		PasswordMd5: req.Password,
		CreateTime:  time.Now(),
	}
}
