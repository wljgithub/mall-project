package dto

type Address struct {
	AddressId     int    `json:"addressId"`
	CityName      string `json:"cityName"`
	DefaultFlag   int    `json:"defaultFlag"`
	DetailAddress string `json:"detailAddress"`
	ProvinceName  string `json:"provinceName"`
	RegionName    string `json:"regionName"`
	UserId        int    `json:"userId"`
	UserName      string `json:"userName"`
	UserPhone     string `json:"userPhone"`
}

type GetAddrRsp struct {
	Address
}

type CreateAddressReq struct {
	CityName      string `json:"cityName"`
	DefaultFlag   int    `json:"defaultFlag"`
	DetailAddress string `json:"detailAddress"`
	ProvinceName  string `json:"provinceName"`
	RegionName    string `json:"regionName"`
	UserName      string `json:"userName"`
	UserPhone     string `json:"userPhone"`
}
type UpdateAddressReq struct {
	Address
}

// route: /api/v1/address/:addressId
type GetAddressDetailRsp struct {
	Address
}

// GET: /api/v1/address/default
type GetDefaultAddressRsp struct {
	Address
}
