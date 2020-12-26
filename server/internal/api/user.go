package api

import "github.com/gin-gonic/gin"

func (this *HttpServer)SayHello(c *gin.Context)  {
	c.JSON(200,"hello")
}
