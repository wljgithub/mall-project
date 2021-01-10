package mapper

import (
	"github.com/wljgithub/mall-project/internal/dto"
	"github.com/wljgithub/mall-project/internal/model"
	"strconv"
	"time"
)

func OrderModelToGetOrderDto(order *model.Order) *dto.GetOrderRsp {
	goods := make([]dto.OrderGoodsItem, len(order.Goods))
	for i := range goods {
		goods[i].GoodsId = order.Goods[i].GoodsId
		goods[i].SellingPrice = order.Goods[i].SellingPrice
		goods[i].GoodsName = order.Goods[i].GoodsName
		goods[i].GoodsCoverImg = order.Goods[i].GoodsCoverImg
		goods[i].GoodsCount = order.Goods[i].GoodsCount
	}
	return &dto.GetOrderRsp{
		CreateTime:             order.CreateTime,
		NewBeeMallOrderItemVOS: goods,
		OrderId:                order.OrderId,
		OrderNo:                order.OrderNo,
		OrderStatus:            order.OrderStatus,
		OrderStatusString:      strconv.Itoa(order.OrderStatus),
		PayType:                order.PayType,
		TotalPrice:             order.TotalPrice,
	}
}
func OrderListToOrderIds(orderList []model.Order) []int {
	orderIds := make([]int, len(orderList))
	for i, order := range orderList {
		orderIds[i] = order.OrderId
	}
	return orderIds
}
func ConcatOrderAndGoodsItem(orderList []model.Order, goodsItem []model.GoodsItem, totalPage int) *dto.GetOrderListRsp {
	m := make(map[int][]dto.OrderGoodsItem)
	for _, item := range goodsItem {
		dtoGoodsItem := dto.OrderGoodsItem{
			GoodsCount:    item.GoodsCount,
			GoodsCoverImg: item.GoodsCoverImg,
			GoodsId:       item.GoodsId,
			GoodsName:     item.GoodsName,
			SellingPrice:  item.SellingPrice,
		}
		m[item.OrderId] = append(m[item.OrderId], dtoGoodsItem)
	}

	dtoOrderList := make([]dto.Order, len(orderList))
	for i, order := range orderList {
		dtoOrderList[i].NewBeeMallOrderItemVOS = m[order.OrderId]
		dtoOrderList[i].OrderId = order.OrderId
		dtoOrderList[i].TotalPrice = order.TotalPrice
		dtoOrderList[i].PayType = order.PayType
		dtoOrderList[i].OrderStatus = order.OrderStatus
		dtoOrderList[i].OrderNo = order.OrderNo
		dtoOrderList[i].CreateTime = order.CreateTime
		dtoOrderList[i].OrderStatusString = OrderStatusMapping(order.OrderStatus)
	}
	return &dto.GetOrderListRsp{List: dtoOrderList, TotalPage: totalPage}
}
func GenerateOrder(uid int, orderNo string, goodsItems []model.PsuedoOrderItemModel) model.Order {
	var price int
	for _, item := range goodsItems {
		price += item.SellingPrice * item.GoodsCount
	}
	now := time.Now()
	return model.Order{
		OrderNo:    orderNo,
		UserId:     uid,
		TotalPrice: price,
		CreateTime: now,
		UpdateTime: now,
		PayTime:    now,
	}
}
func GenerateOrderGoods(orderId int, cartGoodsItems []model.PsuedoOrderItemModel) []model.GoodsItem {
	goods := make([]model.GoodsItem, len(cartGoodsItems))
	now := time.Now()
	for i, cartItem := range cartGoodsItems {
		goods[i].OrderId = orderId
		goods[i].GoodsId = cartItem.GoodsId
		goods[i].GoodsName = cartItem.GoodsName
		goods[i].GoodsCount = cartItem.GoodsCount
		goods[i].SellingPrice = cartItem.SellingPrice
		goods[i].CreateTime = now
		goods[i].GoodsCoverImg = cartItem.GoodsCoverImg
	}
	return goods
}
func OrderStatusMapping(status int) string {
	var statusMapping string
	switch status {
	case 1:
		statusMapping = "待付款"
	case 2:
		statusMapping = "待发货"
	case 3:
		statusMapping = "待收货"
	case 4:
		statusMapping = "待评价"
	case 5:
		statusMapping = "已取消"
	default:
		statusMapping = "未知状态"
	}
	return statusMapping
}
