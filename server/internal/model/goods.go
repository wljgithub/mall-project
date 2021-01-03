package model

import "time"

type Goods struct {
	GoodsId            int
	GoodsName          string
	GoodsIntro         string
	GoodsCategoryId    int
	GoodsCoverImg      string
	GoodsCarousel      string
	GoodsDetailContent string
	OriginalPrice      int
	SellingPrice       int
	StockNum           int
	Tag                string
	GoodsSellStatus    int
	CreateUser         int
	CreateTime         time.Time
	UpdateUser         int
	UpdateTime         time.Time
}

func (Goods) TableName() string {
	return "tb_newbee_mall_goods_info"
}
