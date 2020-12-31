package repository

import (
	"context"
	xerrors "github.com/pkg/errors"
	"github.com/wljgithub/mall-project/internal/dto"
	"github.com/wljgithub/mall-project/internal/model"
	"gorm.io/gorm"
)

type CartRepo interface {
	GetCartList(ctx context.Context, uid int, offset, limit int) ([]dto.CartItemInfo, error)
	CountCartAmount(ctx context.Context, uid int) (int, error)
	CreateCartItem(ctx context.Context, item model.Cart) error
	DeleteCartItem(ctx context.Context, cid string, uid string) error
	UpdateCartItem(cart model.Cart) error
	BatchGetCartItem(ctx context.Context, ids []string) ([]dto.CartItemInfo, error)
}

func (this *Repo) GetCartList(ctx context.Context, uid int, offset, limit int) ([]dto.CartItemInfo, error) {
	cartList := make([]dto.CartItemInfo, 0)
	err := this.db.Table("mall.tb_newbee_mall_shopping_cart_item  cart").
		Select(`cart.cart_item_id ,
					cart.goods_id ,
					cart.goods_count ,
					goods.goods_name ,
					goods.goods_cover_img ,
					goods.selling_price `).
		Joins("inner join mall.tb_newbee_mall_goods_info  goods on goods.goods_id = cart.goods_id").
		Where("cart.user_id = ?", uid).
		Limit(limit).
		Offset(offset).
		Find(&cartList).Error
	return cartList, xerrors.Wrapf(err, "")

}
func (this *Repo) CountCartAmount(ctx context.Context, uid int) (int, error) {
	var amount int64
	err := this.db.Model(&model.Cart{}).Where("user_id = ?", uid).Select("count(*)").Count(&amount).Error
	return int(amount), xerrors.Wrapf(err, "")
}
func (this *Repo) CreateCartItem(ctx context.Context, item model.Cart) error {
	cart := &model.Cart{}
	err := this.db.
		Where(" goods_id = ? AND user_id = ?", item.GoodsId, item.UserId).
		First(cart).Error
	if err == nil {
		return xerrors.Wrapf(this.db.Model(&model.Cart{}).
			Where(" goods_id = ? AND user_id = ?", item.GoodsId, item.UserId).
			Update("goods_count", gorm.Expr("goods_count + ?", item.GoodsCount)).Error, "")
	} else if err == gorm.ErrRecordNotFound {
		return xerrors.Wrapf(this.db.Create(&item).Error, "")
	}
	return err
}
func (this *Repo) DeleteCartItem(ctx context.Context, cid string, uid string) error {
	err := this.db.Model(&model.Cart{}).Where("user_id = ? AND cart_item_id = ?", uid, cid).Delete(&model.Cart{}).Error
	return xerrors.Wrapf(err, "")

}
func (this *Repo) UpdateCartItem(cart model.Cart) error {
	err := this.db.Model(&model.Cart{}).Where("cart_item_id = ?", cart.CartItemId).Updates(cart).Error
	return xerrors.Wrapf(err, "")
}
func (this *Repo) BatchGetCartItem(ctx context.Context, ids []string) ([]dto.CartItemInfo, error) {
	cartList := make([]dto.CartItemInfo, 0)
	err := this.db.Table("mall.tb_newbee_mall_shopping_cart_item  cart").
		Select(`cart.cart_item_id ,
					cart.goods_id ,
					cart.goods_count ,
					goods.goods_name ,
					goods.goods_cover_img ,
					goods.selling_price `).
		Joins("inner join mall.tb_newbee_mall_goods_info  goods on goods.goods_id = cart.goods_id").
		Where("cart.cart_item_id IN ?", ids).
		Find(&cartList).Error
	return cartList, xerrors.Wrapf(err, "")
}
