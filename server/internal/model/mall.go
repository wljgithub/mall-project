package model

import "time"

type Carousel struct {
	CarouselId   int
	CarouselUrl  string
	RedirectUrl  string
	CarouselRank int
	IsDeleted    int
	CreateTime   time.Time
	CreateUser   int
	UpdateTime   time.Time
	UpdateUser   int
}

func (Carousel) TableName() string {
	return "tb_newbee_mall_carousel"
}

type MallCategory struct {
	CategoryId    int
	CategoryLevel int
	ParentId      int
	CategoryName  string
	CategoryRank  int
	IsDeleted     int
	CreateTime    time.Time
	CreateUser    int
	UpdateTime    time.Time
	UpdateUser    int
}

func (MallCategory) TableName() string {
	return "tb_newbee_mall_goods_category"
}
