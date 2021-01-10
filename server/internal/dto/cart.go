package dto

type CartItemInfo struct {
	CartItemId    int    `json:"cartItemId" `
	GoodsCount    int    `json:"goodsCount" `
	GoodsCoverImg string `json:"goodsCoverImg" `
	GoodsId       int    `json:"goodsId" `
	GoodsName     string `json:"goodsName" `
	SellingPrice  int    `json:"sellingPrice" `
}

// GET: /api/v1/shop-cart
type GetCartListReq struct {
	Page     int `form:"page"`
	PageSize int `form:"pageSize"`
}

type GetCartListRsp struct {
	CartList    []CartItemInfo
	TotalAmount int
}

// POST: /api/v1/shop-cart
type CreateCartItemReq struct {
	GoodsCount int `json:"goodsCount"`
	GoodsId    int `json:"goodsId"`
}

// PUT: /api/v1/shop-cart
type UpdateCartItemReq struct {
	CartItemId int `json:"cartItemId"`
	GoodsCount int `json:"goodsCount"`
}
