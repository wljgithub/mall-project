package api

import (
	"github.com/gin-gonic/gin"
	"github.com/wljgithub/mall-project/internal/dto"
	"github.com/wljgithub/mall-project/pkg/errno"
	"github.com/wljgithub/mall-project/pkg/handler"
	"github.com/wljgithub/mall-project/pkg/log"
)

func (this *HttpServer) GetUserInfo(c *gin.Context) {
	uid :=c.GetInt("uid")
	if uid <=0{
		log.Info("failed to parse uid")
		handler.SendResponse(c, errno.ErrInternalServerError, nil)
		return
	}
	userInfo,err:=this.srv.GetUserInfo(uid)
	if err != nil {
		log.Infof("getUserInfo service err: %+v", err)
		handler.SendResponse(c, err, nil)
		return
	}
	handler.SendResponse(c, errno.OK, userInfo)
}

func (this *HttpServer) Login(c *gin.Context) {
	var req dto.LoginReq
	if err := c.ShouldBind(&req); err != nil {
		log.Infof("login bind err:%+v", err)
		handler.SendResponse(c, errno.ErrBind, nil)
		return
	}

	token, err := this.srv.Login(req)
	if err != nil {
		log.Infof("login service err: %+v", err)
		handler.SendResponse(c, err, nil)
		return
	}
	handler.SendResponse(c, errno.OK, token)
}

func (this *HttpServer) Register(c *gin.Context) {

	var req dto.RegisterReq
	if err := c.ShouldBind(&req); err != nil {
		log.Infof("register bind err:%+v", err)
		handler.SendResponse(c, errno.ErrBind, nil)
		return
	}

	err := this.srv.Register(req)
	if err != nil {
		log.Infof("register service err: %+v", err)
		handler.SendResponse(c, err, nil)
		return
	}
	handler.SendResponse(c, errno.OK, nil)
}

func (this *HttpServer) UpdateUserInfo(c *gin.Context) {
	var req dto.UpdateUserInfoReq
	if err := c.ShouldBind(&req); err != nil {
		log.Infof("updateUserInfo bind err:%+v", err)
		handler.SendResponse(c, errno.ErrBind, nil)
		return
	}
	uid := c.GetInt("uid")
	err := this.srv.UpdateUserInfo(uid,req)
	if err != nil {
		log.Infof("login service err: %+v", err)
		handler.SendResponse(c, err, nil)
		return
	}
	handler.SendResponse(c, errno.OK, nil)
}

func (this *HttpServer) Logout(c *gin.Context) {
	uid :=c.GetInt("uid")
	if uid <=0{
		log.Info("logout failed to parse uid")
		handler.SendResponse(c, errno.ErrInternalServerError, nil)
		return
	}
	err:=this.srv.Logout(uid)
	if err != nil {
		log.Infof("logout service err: %+v", err)
		handler.SendResponse(c, err, nil)
		return
	}
	handler.SendResponse(c, errno.OK, nil)
}
