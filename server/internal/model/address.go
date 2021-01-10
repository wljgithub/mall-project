package model

import "time"

type Address struct {
	AddressId     int
	UserId        int
	UserName      string
	UserPhone     string
	DefaultFlag   int
	ProvinceName  string
	CityName      string
	RegionName    string
	DetailAddress string
	IsDeleted     int
	CreateTime    time.Time
	UpdateTime    time.Time
}

func (Address) TableName() string {
	return "tb_newbee_mall_user_address"
}
