package api

import (
	"github.com/gin-gonic/gin"
	"github.com/wljgithub/mall-project/internal/api/middleware"
)

func (this *HttpServer) Load(eg *gin.Engine, mw ...gin.HandlerFunc) *gin.Engine {
	eg.Use(middleware.Options())
	eg.Use(middleware.Secure())
	eg.Use(middleware.NoCache())
	eg.Use(mw...)

	apiV1 := eg.Group("api/v1")

	apiV1.GET("/index-infos", this.MallIndex)
	apiV1.GET("/categories", this.GetCategories)

	// user handler
	user := apiV1.Group("/user")
	{
		user.POST("/login", this.Login)
		user.POST("/register", this.Register)
	}
	userWithPermission := apiV1.Group("/user", middleware.AuthMiddleware())
	{
		userWithPermission.GET("/info", this.GetUserInfo)
		userWithPermission.PUT("/info", this.UpdateUserInfo)
		userWithPermission.POST("/logout", this.Logout)
	}

	addr := apiV1.Group("/address", middleware.AuthMiddleware())
	{
		addr.GET("/", this.GetAddrList)
		addr.POST("/", this.CreateAddress)
		addr.PUT("/", this.UpdateAddress)
		addr.GET("/:addressId", this.GetAddressDetail)
		addr.DELETE("/:addressId", this.DeleteAddress)
		addr.GET("/:addressId/default",this.GetDefaultAddress)
	}
	return eg
}
