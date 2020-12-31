package mapper

import (
	"github.com/wljgithub/mall-project/internal/dto"
	"github.com/wljgithub/mall-project/internal/model"
	"time"
)

func CreateCartItemToCartModel(uid int, req dto.CreateCartItemReq) model.Cart {
	return model.Cart{
		UserId:     uid,
		GoodsId:    req.GoodsId,
		GoodsCount: req.GoodsCount,
		IsDeleted:  0,
		CreateTime: time.Now(),
		UpdateTime: time.Now(),
	}
}
func UpdateCartDtoToCartModel(req dto.UpdateCartItemReq)model.Cart  {
	return model.Cart{CartItemId: req.CartItemId,GoodsCount: req.GoodsCount}
}