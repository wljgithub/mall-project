package dto

import "time"

type Order struct {
	CreateTime             time.Time        `json:"createTime"`
	NewBeeMallOrderItemVOS []OrderGoodsItem `json:"newBeeMallOrderItemVOS"`
	OrderId                int              `json:"orderId"`
	OrderNo                string           `json:"orderNo"`
	OrderStatus            int              `json:"orderStatus"`
	OrderStatusString      string           `json:"orderStatusString"`
	PayType                int              `json:"payType"`
	TotalPrice             int              `json:"totalPrice"`
}
type OrderGoodsItem struct {
	GoodsCount    int    `json:"goodsCount"`
	GoodsCoverImg string `json:"goodsCoverImg"`
	GoodsId       int    `json:"goodsId"`
	GoodsName     string `json:"goodsName"`
	SellingPrice  int    `json:"sellingPrice"`
}

// GET: /api/v1/order
type GetOrderListReq struct {
	PageNumber int `form:"pageNumber"`
	PageSize   int `form:"pageSize"`
	Status     int `form:"status"`
}

type GetOrderListRsp struct {
	List []Order `json:"list"`
}

// GET: /api/v1/order/{orderNo}
type GetOrderRsp = Order

// Get: /api/v1/paySuccess
type PaySuccessReq struct {
	OrderNo string `form:"orderNo"`
	PayType int    `form:"payType"`
}

// POST: /api/v1/saveOrder
type PlaceOrderReq struct {
	AddressId   int   `json:"addressId"`
	CartItemIds []int `json:"cartItemIds"`
}
