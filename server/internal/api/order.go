package api

import (
	"github.com/gin-gonic/gin"
	"github.com/wljgithub/mall-project/internal/dto"
	"github.com/wljgithub/mall-project/pkg/errno"
	"github.com/wljgithub/mall-project/pkg/handler"
	"github.com/wljgithub/mall-project/pkg/log"
)

func (this *HttpServer) GetOrderList(c *gin.Context) {
	uid := c.GetInt("uid")
	if uid <= 0 {
		log.Info("failed to parse uid in getOrderList")
		handler.SendResponse(c, errno.ErrInternalServerError, nil)
		return
	}
	var req dto.GetOrderListReq
	if err := c.ShouldBind(&req); err != nil {
		log.Infof("getOrderList bind err:%+v", err)
		handler.SendResponse(c, errno.ErrBind, nil)
		return
	}

	orderList, err := this.srv.GetOrderList(uid, req)
	if err != nil {
		log.Infof("getOrderList service err: %+v", err)
		handler.SendResponse(c, err, nil)
		return
	}
	handler.SendResponse(c, nil, orderList)
}
func (this *HttpServer) GetOrder(c *gin.Context) {
	uid := c.GetInt("uid")
	if uid <= 0 {
		log.Info("failed to parse uid in getOrder")
		handler.SendResponse(c, errno.ErrInternalServerError, nil)
		return
	}
	orderNo := c.Param("orderNo")

	order, err := this.srv.GetOrder(uid, orderNo)
	if err != nil {
		log.Infof("getOrder service err: %+v", err)
		handler.SendResponse(c, err, nil)
		return
	}
	handler.SendResponse(c, nil, order)
}
func (this *HttpServer) CancelOrder(c *gin.Context) {
	uid := c.GetInt("uid")
	if uid <= 0 {
		log.Info("failed to parse uid in cancelOrder")
		handler.SendResponse(c, errno.ErrInternalServerError, nil)
		return
	}
	orderNo := c.Param("orderNo")
	err := this.srv.CancelOrder(uid, orderNo)
	if err != nil {
		log.Infof("cancelOrder service err: %+v", err)
		handler.SendResponse(c, err, nil)
		return
	}
	handler.SendResponse(c, nil, nil)
}
func (this *HttpServer) FinishOrder(c *gin.Context) {
	uid := c.GetInt("uid")
	if uid <= 0 {
		log.Info("failed to parse uid in finishOrder")
		handler.SendResponse(c, errno.ErrInternalServerError, nil)
		return
	}
	orderNo := c.Param("orderNo")
	err := this.srv.FinishOrder(uid, orderNo)
	if err != nil {
		log.Infof("finishOrder service err: %+v", err)
		handler.SendResponse(c, err, nil)
		return
	}
	handler.SendResponse(c, nil, nil)
}
func (this *HttpServer) PayOrder(c *gin.Context) {
	uid := c.GetInt("uid")
	if uid <= 0 {
		log.Info("failed to parse uid in payOrder")
		handler.SendResponse(c, errno.ErrInternalServerError, nil)
		return
	}
	var req dto.PaySuccessReq
	if err := c.ShouldBind(&req); err != nil {
		log.Infof("PayOrder bind err:%+v", err)
		handler.SendResponse(c, errno.ErrBind, nil)
		return
	}

	err := this.srv.PaySuccess(uid, req)
	if err != nil {
		log.Infof("payOrder service err: %+v", err)
		handler.SendResponse(c, err, nil)
		return
	}
	handler.SendResponse(c, nil, nil)
}
func (this *HttpServer) PlaceOrder(c *gin.Context) {
	uid := c.GetInt("uid")
	if uid <= 0 {
		log.Info("failed to parse uid in placeOrder")
		handler.SendResponse(c, errno.ErrInternalServerError, nil)
		return
	}
	var req dto.PlaceOrderReq
	if err := c.ShouldBind(&req); err != nil {
		log.Infof("placeOrder bind err:%+v", err)
		handler.SendResponse(c, errno.ErrBind, nil)
		return
	}

	orderNo,err := this.srv.PlaceOrder(uid, req)
	if err != nil {
		log.Infof("placeOrder service err: %+v", err)
		handler.SendResponse(c, err, nil)
		return
	}
	handler.SendResponse(c, nil, orderNo)
}
