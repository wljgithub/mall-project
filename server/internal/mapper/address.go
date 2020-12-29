package mapper

import (
	"github.com/wljgithub/mall-project/internal/dto"
	"github.com/wljgithub/mall-project/internal/model"
	"time"
)

func AddressModelToDtoGetAddrRsp(address []model.Address) []dto.GetAddrRsp {
	result := make([]dto.GetAddrRsp, len(address))
	for i, addr := range address {
		result[i].UserId = addr.UserId
		result[i].UserName = addr.UserName
		result[i].AddressId = addr.AddressId
		result[i].DetailAddress = addr.DetailAddress
		result[i].CityName = addr.CityName
		result[i].DefaultFlag = addr.DefaultFlag
		result[i].ProvinceName = addr.ProvinceName
		result[i].RegionName = addr.RegionName
		result[i].UserPhone = addr.UserPhone
	}
	return result
}
func CreateAddressDtoToAddressModel(addr dto.CreateAddressReq, uid int) model.Address {
	return model.Address{
		UserId:        uid,
		UserName:      addr.UserName,
		UserPhone:     addr.UserPhone,
		DefaultFlag:   addr.DefaultFlag,
		ProvinceName:  addr.ProvinceName,
		CityName:      addr.CityName,
		RegionName:    addr.RegionName,
		DetailAddress: addr.DetailAddress,
		CreateTime:    time.Now(),
		UpdateTime:    time.Now(),
	}
}
func UpdateAddressDtoToAddressModel(addr dto.UpdateAddressReq, uid int) model.Address {
	return model.Address{
		AddressId:     addr.AddressId,
		UserId:        addr.UserId,
		UserName:      addr.UserName,
		UserPhone:     addr.UserPhone,
		DefaultFlag:   addr.DefaultFlag,
		ProvinceName:  addr.ProvinceName,
		CityName:      addr.CityName,
		RegionName:    addr.RegionName,
		DetailAddress: addr.DetailAddress,
		UpdateTime:    time.Now(),
	}
}
func AddressModelToGetAdressDetailDto(addr model.Address) *dto.GetAddressDetailRsp {
	return &dto.GetAddressDetailRsp{dto.Address{
		AddressId:     addr.AddressId,
		CityName:      addr.CityName,
		DefaultFlag:   addr.DefaultFlag,
		DetailAddress: addr.DetailAddress,
		ProvinceName:  addr.ProvinceName,
		RegionName:    addr.RegionName,
		UserId:        addr.UserId,
		UserName:      addr.UserName,
		UserPhone:     addr.UserPhone,
	}}
}
func AddressModelToDefaultAddressDto(addr model.Address) *dto.GetDefaultAddressRsp {
	return &dto.GetDefaultAddressRsp{dto.Address{
		AddressId:     addr.AddressId,
		CityName:      addr.CityName,
		DefaultFlag:   addr.DefaultFlag,
		DetailAddress: addr.DetailAddress,
		ProvinceName:  addr.ProvinceName,
		RegionName:    addr.RegionName,
		UserId:        addr.UserId,
		UserName:      addr.UserName,
		UserPhone:     addr.UserPhone,
	}}
}
