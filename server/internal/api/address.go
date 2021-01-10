package api

import (
	"github.com/gin-gonic/gin"
	"github.com/wljgithub/mall-project/internal/dto"
	"github.com/wljgithub/mall-project/pkg/errno"
	"github.com/wljgithub/mall-project/pkg/handler"
	"github.com/wljgithub/mall-project/pkg/log"
)

func (this *HttpServer) GetAddrList(c *gin.Context) {
	uid := c.GetInt("uid")
	if uid <= 0 {
		log.Info("failed to parse uid in getAddrList")
		handler.SendResponse(c, errno.ErrInternalServerError, nil)
		return
	}

	addrList, err := this.srv.GetAddrList(uid)
	if err != nil {
		log.Infof("getAddrList service err:%+v", err)
		handler.SendResponse(c, err, nil)
	}
	handler.SendResponse(c, nil, addrList)
}
func (this *HttpServer) CreateAddress(c *gin.Context) {
	var req dto.CreateAddressReq
	if err := c.ShouldBind(&req); err != nil {
		log.Infof("createAddress bind err:%+v", err)
		handler.SendResponse(c, errno.ErrBind, nil)
		return
	}
	uid := c.GetInt("uid")
	err := this.srv.CreateAddress(req, uid)
	if err != nil {
		log.Infof("createAddress service err: %+v", err)
		handler.SendResponse(c, err, nil)
		return
	}
	handler.SendResponse(c, nil, nil)
}

func (this *HttpServer) UpdateAddress(c *gin.Context) {
	var req dto.UpdateAddressReq
	if err := c.ShouldBind(&req); err != nil {
		log.Infof("updateAddress bind err:%+v", err)
		handler.SendResponse(c, errno.ErrBind, nil)
		return
	}
	uid := c.GetInt("uid")
	err := this.srv.UpdateAddress(req, uid)
	if err != nil {
		log.Infof("updateAddress service err: %+v", err)
		handler.SendResponse(c, err, nil)
		return
	}
	handler.SendResponse(c, nil, nil)
}
func (this *HttpServer) GetAddressDetail(c *gin.Context) {
	addressId := c.Param("addressId")
	address, err := this.srv.GetAddressDetail(addressId)
	if err != nil {
		log.Infof("getAddressDetail service err: %+v", err)
		handler.SendResponse(c, err, nil)
		return
	}
	handler.SendResponse(c, nil, address)
}
func (this *HttpServer) DeleteAddress(c *gin.Context) {

	addressId := c.Param("addressId")
	err := this.srv.DeleteAdress(addressId)
	if err != nil {
		log.Infof("deleteAddress service err: %+v", err)
		handler.SendResponse(c, err, nil)
		return
	}
	handler.SendResponse(c, nil, nil)
}
func (this *HttpServer) GetDefaultAddress(c *gin.Context) {
	address, err := this.srv.GetDefaultAddress()
	if err != nil {
		log.Infof("getDefaultAddress service err: %+v", err)
		handler.SendResponse(c, err, nil)
		return
	}
	handler.SendResponse(c, nil, address)
}
