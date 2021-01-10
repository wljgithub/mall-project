package model

import "time"

type Cart struct {
	CartItemId int
	UserId     int
	GoodsId    int
	GoodsCount int
	IsDeleted  int
	CreateTime time.Time
	UpdateTime time.Time
}
func (Cart) TableName() string {
	return "tb_newbee_mall_shopping_cart_item"
}
