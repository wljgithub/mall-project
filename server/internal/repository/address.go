package repository

import (
	"context"
	xerrors "github.com/pkg/errors"
	"github.com/wljgithub/mall-project/internal/model"
)

type AddressRepo interface {
	FetchAddressList(ctx context.Context, uid int) ([]model.Address, error)
	CreateAddress(ctx context.Context, address model.Address) error
	UpdateAddress(ctx context.Context, address model.Address) error
	GetAddressDetail(addrId string) (model.Address, error)
	DeleteAddress(ctx context.Context, id string) error
	GetDefaultAddress(ctx context.Context) (model.Address, error)
}

func (this *Repo) FetchAddressList(ctx context.Context, uid int) ([]model.Address, error) {
	address := make([]model.Address, 0)
	err := this.db.Model(&model.Address{}).Find(&address).Error
	if err != nil {
		return nil, xerrors.Wrapf(err, "")
	}
	return address, nil
}
func (this *Repo) CreateAddress(ctx context.Context, address model.Address) error {
	err := this.db.Create(&address).Error
	return xerrors.Wrapf(err, "")
}
func (this *Repo) UpdateAddress(ctx context.Context, address model.Address) error {
	err := this.db.Model(&model.Address{}).Where("address_id = ?", address.AddressId).Updates(&address).Error
	return xerrors.Wrapf(err, "")
}
func (this *Repo) GetAddressDetail(addrId string) (model.Address, error) {
	addr := model.Address{}
	err := this.db.Model(&model.Address{}).Where("address_id = ?", addrId).First(&addr).Error
	return addr, xerrors.Wrapf(err, "")
}
func (this *Repo) DeleteAddress(ctx context.Context, id string) error {
	err := this.db.Model(&model.Address{}).Where("address_id = ?", id).Delete(&model.Address{}).Error
	return xerrors.Wrapf(err, "")
}
func (this *Repo)  GetDefaultAddress(ctx context.Context) (model.Address, error){
	address:=model.Address{}
	err:=this.db.Model(&model.Address{}).Where("default_flag = ?",1).First(&address).Error
	return address,xerrors.Wrapf(err,"")
}