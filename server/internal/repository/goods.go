package repository

import (
	"context"
	"fmt"
	xerrors "github.com/pkg/errors"
	"github.com/wljgithub/mall-project/internal/dto"
	"github.com/wljgithub/mall-project/internal/model"
	"gorm.io/gorm"
)

type GoodsRepo interface {
	GetGoods(ctx context.Context, id string) (model.Goods, error)
	SearchGoodsByCaregory(ctx context.Context, id string, by string, offset int, limit int) ([]dto.SearchGoodsItem, error)
	SearchGoodsByName(ctx context.Context, keyword string, by string, offset int, limit int) ([]dto.SearchGoodsItem, error)
	CountGoodsByCategory(ctx context.Context, categoryId string) (int, error)
	CountGoodsByName(ctx context.Context, keyword string) (int, error)
}

func (this *Repo) GetGoods(ctx context.Context, id string) (model.Goods, error) {
	goods := model.Goods{}
	err := this.db.Model(&model.Goods{}).Where("goods_id = ?", id).First(&goods).Error
	return goods, xerrors.Wrapf(err, "")
}
func (this *Repo) searchGoods() *gorm.DB {
	//cartList := make([]dto.CartItemInfo, 0)
	return this.db.Model(&model.Goods{}).
		Select(`
					goods_id ,
					goods_intro
					goods_name ,
					goods_cover_img ,
					selling_price `)
	//Joins("inner join mall.tb_newbee_mall_goods_info  goods on goods.goods_id = cart.goods_id")
}
func (this *Repo) searchGoodsOrder(sql *gorm.DB, by string) *gorm.DB {
	switch by {
	case "new":
		sql = sql.Order("update_time DESC")
	case "price":
		sql = sql.Order("selling_price DESC")
	}
	return sql
}
func (this *Repo) SearchGoodsByCaregory(ctx context.Context, id string, by string, offset int, limit int) ([]dto.SearchGoodsItem, error) {
	goods := []dto.SearchGoodsItem{}
	err := this.searchGoodsOrder(this.searchGoods().Where("goods_category_id = ?", id), by).
		Offset(offset).
		Limit(limit).Find(&goods).Error
	return goods, xerrors.Wrapf(err, "")
}
func (this *Repo) CountGoodsByCategory(ctx context.Context, categoryId string) (int, error) {
	var count int64
	err := this.db.Model(&model.Goods{}).Where("goods_category_id = ?", categoryId).Count(&count).Error
	return int(count), xerrors.Wrapf(err, "")
}
func (this *Repo) CountGoodsByName(ctx context.Context, keyword string) (int, error) {
	var count int64
	err := this.db.Model(&model.Goods{}).
		Where("goods_name LIKE ?", fmt.Sprintf("%%%s%%", keyword)).
		Count(&count).Error
	return int(count), xerrors.Wrapf(err, "")
}
func (this *Repo) SearchGoodsByName(ctx context.Context, keyword string, by string, offset int, limit int) ([]dto.SearchGoodsItem, error) {
	goods := []dto.SearchGoodsItem{}
	err := this.searchGoodsOrder(this.searchGoods().
		Where("goods_name LIKE ?", fmt.Sprintf("%%%s%%", keyword)), by).
		Offset(offset).
		Limit(limit).Find(&goods).Error
	return goods, xerrors.Wrapf(err, "")
}
