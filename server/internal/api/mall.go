package api

import (
	"github.com/gin-gonic/gin"
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
