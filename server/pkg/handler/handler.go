package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/wljgithub/mall-project/pkg/errno"
)

type Response struct {
	Code    int         `json:"resultCode"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func SendResponse(c *gin.Context, err error, data interface{}) {

	statusCode, serviceCode, message := errno.DecodeErr(err)

	c.JSON(statusCode, Response{
		Code:    serviceCode,
		Message: message,
		Data:    data,
	})
}
