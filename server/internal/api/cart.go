package api

import (
	"github.com/gin-gonic/gin"
	"github.com/wljgithub/mall-project/internal/dto"
	"github.com/wljgithub/mall-project/pkg/errno"
	"github.com/wljgithub/mall-project/pkg/handler"
	"github.com/wljgithub/mall-project/pkg/log"
	"strconv"
	"strings"
)

func (this *HttpServer) GetCartList(c *gin.Context) {
	uid := c.GetInt("uid")
	if uid <= 0 {
		log.Info("failed to parse uid in getCartList")
		handler.SendResponse(c, errno.ErrInternalServerError, nil)
		return
	}
	var req dto.GetCartListReq
	if err := c.ShouldBind(&req); err != nil {
		log.Infof("getCartList bind err:%+v", err)
		handler.SendResponse(c, errno.ErrBind, nil)
		return
	}
	cartList, err := this.srv.GetCartList(uid, req)
	if err != nil {
		log.Infof("getCartList service err: %+v", err)
		handler.SendResponse(c, err, nil)
		return
	}
	handler.SendResponse(c, nil, cartList)
}
func (this *HttpServer) CreateCartItem(c *gin.Context) {
	uid := c.GetInt("uid")
	if uid <= 0 {
		log.Info("failed to parse uid in createCartItem")
		handler.SendResponse(c, errno.ErrInternalServerError, nil)
		return
	}

	var req dto.CreateCartItemReq
	if err := c.ShouldBind(&req); err != nil {
		log.Infof("createCartItem bind err:%+v", err)
		handler.SendResponse(c, errno.ErrBind, nil)
		return
	}
	err := this.srv.CreateCartItem(uid, req)
	if err != nil {
		log.Infof("createCartItem service err: %+v", err)
		handler.SendResponse(c, err, nil)
		return
	}
	handler.SendResponse(c, nil, nil)
}
func (this *HttpServer) DeleteCartItem(c *gin.Context) {
	uid := c.GetInt("uid")
	if uid <= 0 {
		log.Info("failed to parse uid in getCartList")
		handler.SendResponse(c, errno.ErrInternalServerError, nil)
		return
	}
	cartItemId := c.Param("newBeeMallShoppingCartItemId")
	err := this.srv.DeleteCartItem(cartItemId, strconv.Itoa(uid))
	if err != nil {
		log.Infof("deleteCartItem service err: %+v", err)
		handler.SendResponse(c, err, nil)
		return
	}
	handler.SendResponse(c, nil, nil)
}
func (this *HttpServer) UpdateCartItem(c *gin.Context) {
	var req dto.UpdateCartItemReq
	if err := c.ShouldBind(&req); err != nil {
		log.Infof("login bind err:%+v", err)
		handler.SendResponse(c, errno.ErrBind, nil)
		return
	}

	err := this.srv.UpdateCartItem(req)
	if err != nil {
		log.Infof("login service err: %+v", err)
		handler.SendResponse(c, err, nil)
		return
	}
	handler.SendResponse(c, nil, nil)
}
func (this *HttpServer) BatchGetCartItem(c *gin.Context) {
	cartItemIds := c.QueryArray("cartItemIds")
	if len(cartItemIds) == 0 {
		log.Info("cartItemid is empty in BatchGetCartItem")
		handler.SendResponse(c, errno.ErrBind, nil)
		return
	}
	cartIds := strings.Split(cartItemIds[0], ",")
	cart, err := this.srv.BatchGetCartItem(cartIds)
	if err != nil {
		log.Infof("batchGetCartItem service err: %+v", err)
		handler.SendResponse(c, err, nil)
		return
	}
	handler.SendResponse(c, nil, cart)
}
