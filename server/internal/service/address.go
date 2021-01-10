package service

import (
	"context"
	"github.com/wljgithub/mall-project/internal/dto"
	"github.com/wljgithub/mall-project/internal/mapper"
)

func (this *Service) GetAddrList(uid int) ([]dto.GetAddrRsp, error) {
	address, err := this.Repo.FetchAddressList(context.Background(), uid)
	if err != nil {
		return nil, err
	}
	return mapper.AddressModelToDtoGetAddrRsp(address), nil
}

func (this *Service) GetAddressDetail(addrId string) (*dto.GetAddressDetailRsp, error) {
	addressDetail, err := this.Repo.GetAddressDetail(addrId)
	if err != nil {
		return nil, err
	}
	return mapper.AddressModelToGetAdressDetailDto(addressDetail), nil
}
func (this *Service) CreateAddress(req dto.CreateAddressReq, uid int) error {
	address := mapper.CreateAddressDtoToAddressModel(req, uid)
	return this.Repo.CreateAddress(context.Background(), address)
}
func (this *Service) UpdateAddress(req dto.UpdateAddressReq, uid int) error {
	address := mapper.UpdateAddressDtoToAddressModel(req, uid)
	return this.Repo.UpdateAddress(context.Background(), address)

}
func (this *Service) DeleteAdress(addressId string) error {
	return this.Repo.DeleteAddress(context.Background(), addressId)
}
func (this *Service) GetDefaultAddress() (*dto.GetDefaultAddressRsp, error) {
	address, err := this.Repo.GetDefaultAddress(context.Background())
	if err != nil {
		return nil, err
	}
	return mapper.AddressModelToDefaultAddressDto(address), nil
}
