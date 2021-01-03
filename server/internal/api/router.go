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
	apiV1.GET("/search", this.GoodsSearch)

	// 用户相关接口
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
	// 个人地址相关接口
	addr := apiV1.Group("/address", middleware.AuthMiddleware())
	{
		addr.GET("/", this.GetAddrList)
		addr.POST("/", this.CreateAddress)
		addr.PUT("/", this.UpdateAddress)
		addr.GET("/:addressId", this.GetAddressDetail)
		addr.DELETE("/:addressId", this.DeleteAddress)
		addr.GET("/:addressId/default", this.GetDefaultAddress)
	}

	// 购物车相关接口
	cart := apiV1.Group("/shop-cart", middleware.AuthMiddleware())
	{
		cart.GET("/", this.GetCartList)
		cart.POST("/", this.CreateCartItem)
		cart.DELETE("/:newBeeMallShoppingCartItemId", this.DeleteCartItem)
		cart.PUT("/", this.UpdateCartItem)
		cart.GET("/settle", this.BatchGetCartItem)
	}

	// 订单相关
	order := apiV1.Group("order", middleware.AuthMiddleware())
	{
		order.GET("/", this.GetOrderList)
		order.GET("/:orderNo", this.GetOrder)
		order.PUT("/:orderNo/cancel", this.CancelOrder)
		order.PUT("/:orderNo/finish", this.FinishOrder)
	}
	// 支付mock接口
	payOrder := apiV1.Group("/paySuccess", middleware.AuthMiddleware())
	{
		payOrder.GET("/", this.PayOrder)
	}
	// 生成订单接口
	placeOrder := apiV1.Group("/saveOrder", middleware.AuthMiddleware())
	{
		placeOrder.POST("/", this.PlaceOrder)

	// 商品相关
	goods := apiV1.Group("/goods", middleware.AuthMiddleware())
	{
		goods.GET("/detail/:goodsId", this.GetGoodsDetail)

	}
	return eg
}
