package service

import (
	"context"
	"github.com/wljgithub/mall-project/internal/dto"
	"github.com/wljgithub/mall-project/internal/mapper"
)

func (this *Service) GetCartList(uid int, req dto.GetCartListReq) (*dto.GetCartListRsp, error) {
	var offset, limit int
	if req.Page < 1 {
		req.Page = 1
	}
	offset = (req.Page - 1) * limit
	limit = req.PageSize
	cartList, err := this.Repo.GetCartList(context.Background(), uid, offset, limit)
	if err != nil {
		return nil, err
	}
	totalAmount, err := this.Repo.CountCartAmount(context.Background(), uid)
	if err != nil {
		return nil, err
	}

	return &dto.GetCartListRsp{
		CartList:    cartList,
		TotalAmount: totalAmount,
	}, nil
}
func (this *Service) CreateCartItem(uid int, req dto.CreateCartItemReq) error {
	cartItem := mapper.CreateCartItemToCartModel(uid, req)
	return this.Repo.CreateCartItem(context.Background(), cartItem)
}
func (this *Service) DeleteCartItem(cid string, uid string) error {
	return this.Repo.DeleteCartItem(context.Background(), cid, uid)
}
func (this *Service) UpdateCartItem(req dto.UpdateCartItemReq) error {
	cart := mapper.UpdateCartDtoToCartModel(req)
	return this.Repo.UpdateCartItem(cart)
}
func (this *Service) BatchGetCartItem(ids []string) ([]dto.CartItemInfo, error) {
	return this.Repo.BatchGetCartItem(context.Background(),ids)
}
