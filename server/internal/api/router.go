package api

import "github.com/gin-gonic/gin"

func (this *HttpServer)Load(eg *gin.Engine, mw ...gin.HandlerFunc) *gin.Engine {
	eg.GET("/", this.SayHello)
	return eg
}
