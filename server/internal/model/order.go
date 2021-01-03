package model

import "time"

const (
	NotPay       = 1 + iota // 待付款
	NotDelivered            // 待发货
	NotReceive              //  待收货
	NotComment              //	待评价
	HasCancel               //	已取消
)

const (
	HasPay = 1 + iota
)

type GoodsItem struct {
	OrderItemId   int `gorm:"primaryKey"`
	OrderId       int
	GoodsId       int
	GoodsName     string
	GoodsCoverImg string
	SellingPrice  int
	GoodsCount    int
	CreateTime    time.Time
}

type Order struct {
	OrderId     int `gorm:"primaryKey"`
	OrderNo     string
	UserId      int
	TotalPrice  int
	PayStatus   int
	PayType     int
	PayTime     time.Time
	OrderStatus int
	ExtraInfo   string
	IsDeleted   int
	CreateTime  time.Time
	UpdateTime  time.Time
	Goods       []GoodsItem `gorm:"foreignKey:OrderId"`
}

type PsuedoOrderItemModel struct {
	CartItemId    int
	GoodsId       int
	GoodsCount    int
	GoodsName     string
	GoodsCoverImg string
	SellingPrice  int
}

func (Order) TableName() string {
	return "tb_newbee_mall_order"
}
func (GoodsItem) TableName() string {
	return "tb_newbee_mall_order_item"
}
