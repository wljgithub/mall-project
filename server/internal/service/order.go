package service

import (
	"context"
	"github.com/wljgithub/mall-project/internal/dto"
	"github.com/wljgithub/mall-project/internal/mapper"
)



func (this *Service) GetOrderList(uid int, req dto.GetOrderListReq) (*dto.GetOrderListRsp, error) {
	if req.PageSize < 0 {
		req.PageSize = 5
	}
	if req.PageNumber <= 0 {
		req.PageNumber = 1
	}
	offset := (req.PageNumber - 1) * req.PageSize

	orderList, err := this.Repo.GetOrderList(context.Background(), uid, req.Status, offset, req.PageSize)
	if err != nil {
		return nil, err
	}
	orderIds := mapper.OrderListToOrderIds(orderList)
	goodsItems, err := this.Repo.GetOrderGoodsItems(context.Background(), orderIds)
	if err != nil {
		return nil, err
	}

	return mapper.ConcatOrderAndGoodsItem(orderList, goodsItems), nil

}
func (this *Service) GetOrder(uid int, orderNo string) (*dto.GetOrderRsp, error) {
	order, err := this.Repo.GetOrder(uid, orderNo)
	if err != nil {
		return nil, err
	}
	return mapper.OrderModelToGetOrderDto(order), nil
}
func (this *Service) CancelOrder(uid int, orderNo string) error {
	return this.Repo.CancelOrder(context.Background(), uid, orderNo)
}
func (this *Service) FinishOrder(uid int, orderNo string) error {
	return this.Repo.FinishOrder(context.Background(),uid,orderNo)
}
func (this *Service) PaySuccess(uid int, req dto.PaySuccessReq) error {
	return this.Repo.PaySuccess(context.Background(),uid,req)
}
func (this *Service) PlaceOrder(uid int, req dto.PlaceOrderReq) (string,error) {
	// 根据购物车id查询购物项

	// 生成订单信息
	return this.Repo.PlaceOrder(context.Background(),uid,req)
}
