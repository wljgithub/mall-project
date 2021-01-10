package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/wljgithub/mall-project/pkg/errno"
	"github.com/wljgithub/mall-project/pkg/handler"
	"github.com/wljgithub/mall-project/pkg/token"
)

// AuthMiddleware 认证中间件
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Parse the json web token.
		ctx, err := token.ParseRequest(c)
		//log.Infof("context is: %+v", ctx)

		if err != nil {
			handler.SendResponse(c, errno.ErrTokenInvalid, nil)
			c.Abort()
			return
		}
		if expire, err := token.IsExpire(ctx); expire || err != nil {
			handler.SendResponse(c, errno.ErrTokenExpire, nil)
			c.Abort()
			return
		}
		// set uid to context
		c.Set("uid", int(ctx.UserID))

		c.Next()
	}
}
