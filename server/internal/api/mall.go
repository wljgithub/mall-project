package api

import (
	"github.com/gin-gonic/gin"
	"github.com/wljgithub/mall-project/internal/dto"
	"github.com/wljgithub/mall-project/pkg/errno"
	"github.com/wljgithub/mall-project/pkg/handler"
	"github.com/wljgithub/mall-project/pkg/log"
)

func (this *HttpServer) MallIndex(c *gin.Context) {
	mallIndex, err := this.srv.GetMallIndex()
	if err != nil {
		log.Infof("mallindex service err: %+v", err)
		handler.SendResponse(c, err, nil)
		return
	}
	handler.SendResponse(c, err, mallIndex)
}
func (this *HttpServer) GetCategories(c *gin.Context) {
	category, err := this.srv.GetCategory()

	if err != nil {
		log.Infof("getCategories service err: %+v", err)
		handler.SendResponse(c, err, nil)
		return
	}
	handler.SendResponse(c, err, category.SubCategories)
}
func (this *HttpServer) GetGoodsDetail(c *gin.Context) {
	goodsId := c.Param("goodsId")

	goodsDetail, err := this.srv.GetGoodsDetail(goodsId)
	if err != nil {
		log.Infof("getGoodsDetail service err: %+v", err)
		handler.SendResponse(c, err, nil)
		return
	}
	handler.SendResponse(c, err, goodsDetail)
}
func (this *HttpServer) GoodsSearch(c *gin.Context) {
	var req dto.GoodsSearchReq
	if err := c.ShouldBind(&req); err != nil {
		log.Infof("goodsSearch bind err:%+v", err)
		handler.SendResponse(c, errno.ErrBind, nil)
		return
	}

	goods, err := this.srv.GoodsSearch(req)
	if err != nil {
		log.Infof("goodsSearch service err: %+v", err)
		handler.SendResponse(c, err, nil)
		return
	}
	handler.SendResponse(c, nil, goods)
}
