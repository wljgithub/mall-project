package repository

import (
	"context"
	"github.com/google/uuid"
	xerrors "github.com/pkg/errors"
	"github.com/wljgithub/mall-project/internal/dto"
	"github.com/wljgithub/mall-project/internal/mapper"
	"github.com/wljgithub/mall-project/internal/model"
	"gorm.io/gorm"
	"time"
)

type OrderRepo interface {
	GetOrder(uid int, orderNo string) (*model.Order, error)
	GetOrderList(ctx context.Context, uid int, status, offset, limit int) ([]model.Order, error)
	GetOrderGoodsItems(ctx context.Context, ids []int) ([]model.GoodsItem, error)
	CancelOrder(ctx context.Context, uid int, orderNo string) error
	FinishOrder(ctx context.Context, uid int, orderNo string) error
	PaySuccess(ctx context.Context, uid int, req dto.PaySuccessReq) error
	PlaceOrder(ctx context.Context, uid int, req dto.PlaceOrderReq) (string, error)
}

func (this *Repo) GetOrder(uid int, orderNo string) (*model.Order, error) {
	order := model.Order{}
	err := this.db.Where(&model.Order{OrderNo: orderNo, UserId: uid}).First(&order).Error
	if err != nil {
		return &order, xerrors.Wrapf(err, "")
	}
	err = this.db.Where(&model.GoodsItem{OrderId: order.OrderId}).Find(&(order.Goods)).Error
	return &order, xerrors.Wrapf(err, "")
}
func (this *Repo) GetOrderList(ctx context.Context, uid int, status, offset, limit int) ([]model.Order, error) {
	order := []model.Order{}
	sql := this.db.Where(&model.Order{UserId: uid})
	if status != 0 {
		sql = sql.Where(&model.Order{PayStatus: status})
	}
	err := sql.Offset(offset).Limit(limit).Find(&order).Error
	return order, xerrors.Wrapf(err, "")
}
func (this *Repo) GetOrderGoodsItems(ctx context.Context, ids []int) ([]model.GoodsItem, error) {
	goodsItems := make([]model.GoodsItem, 0)
	err := this.db.Model(&model.GoodsItem{}).Where("order_id IN ?", ids).Find(&goodsItems).Error
	return goodsItems, xerrors.Wrapf(err, "")
}
func (this *Repo) CancelOrder(ctx context.Context, uid int, orderNo string) error {
	sql := this.db.Where(&model.Order{UserId: uid, OrderNo: orderNo})
	err := sql.First(&model.Order{}).Error
	if err != nil {
		return err
	}
	return sql.Updates(model.Order{OrderStatus: model.HasCancel}).Error
}
func (this *Repo) FinishOrder(ctx context.Context, uid int, orderNo string) error {
	sql := this.db.Where(&model.Order{UserId: uid, OrderNo: orderNo})
	err := sql.First(&model.Order{}).Error
	if err != nil {
		return err
	}
	return sql.Updates(model.Order{OrderStatus: model.NotComment}).Error
}
func (this *Repo) PaySuccess(ctx context.Context, uid int, req dto.PaySuccessReq) error {
	condition := this.db.Where(&model.Order{UserId: uid, OrderNo: req.OrderNo})

	err := condition.First(&model.Order{}).Error
	if err != nil {
		return err
	}
	err = condition.Updates(&model.Order{
		PayStatus:   model.HasPay,
		PayType:     req.PayType,
		OrderStatus: model.NotDelivered,
		PayTime:     time.Now(),
	}).Error
	return xerrors.Wrapf(err, "")
}
func (this *Repo) PlaceOrder(ctx context.Context, uid int, req dto.PlaceOrderReq) (string, error) {
	var (
		cartItems []model.PsuedoOrderItemModel
		err       error
	)
	// 根据购物车id查询商品信息
	err = this.db.Table("tb_newbee_mall_shopping_cart_item cart").
		Select(`cart.cart_item_id,
						cart.goods_id,
						cart.goods_count,
						goods.goods_name,
						goods.goods_cover_img,
						goods.selling_price`).
		Joins("inner join tb_newbee_mall_goods_info goods on goods.goods_id = cart.goods_id").
		Where("cart_item_id IN ?", req.CartItemIds).Find(&cartItems).Error
	if err != nil {
		return "", xerrors.Wrapf(err, "")
	}
	orderNo := uuid.New().String()
	orderInfo := mapper.GenerateOrder(uid, orderNo, cartItems)

	// 开启事务
	if err := this.db.Transaction(func(tx *gorm.DB) error {
		// 生成订单
		if err := this.db.Create(&orderInfo).Error; err != nil {
			return err
		}
		// 将购物车商品放入订单条目里
		orderGoodsItems := mapper.GenerateOrderGoods(orderInfo.OrderId, cartItems)
		if err := this.db.Create(&orderGoodsItems).Error; err != nil {
			return err
		}
		// 删除购物车中购物项
		if err := this.db.Where("cart_item_id IN ?", req.CartItemIds).Delete(&model.Cart{}).Error; err != nil {
			return err
		}
		return nil
	}); err != nil {
		return "", xerrors.Wrapf(err, "")
	}

	return orderNo, nil
}
