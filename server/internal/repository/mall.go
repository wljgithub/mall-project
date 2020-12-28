package repository

import (
	"context"
	xerrors "github.com/pkg/errors"
	"github.com/wljgithub/mall-project/internal/dto"
	"github.com/wljgithub/mall-project/internal/model"
)

type MallRepo interface {
	FetchAllUser(ctx context.Context) ([]*model.Carousel, error)
	FetchGoodsByType(ctx context.Context, typed int) ([]dto.MallIndexGoods, error)
	FetchAllCategory(ctx context.Context) ([]model.MallCategory, error)
}

func (this *Repo) FetchAllUser(ctx context.Context) ([]*model.Carousel, error) {
	carousels := make([]*model.Carousel, 0)
	err := this.db.Find(&carousels).Error
	if err != nil {
		return nil, xerrors.Wrapf(err, "failed to fetch all user")

	}
	return carousels, nil
}

func (this *Repo) FetchGoodsByType(ctx context.Context, typed int) ([]dto.MallIndexGoods, error) {
	goods := make([]dto.MallIndexGoods, 0)
	err := this.db.Table("tb_newbee_mall_index_config c").
		Select(`g.goods_cover_img,
					g.goods_id,
					g.goods_intro as goods_info,
					g.goods_name,
					g.selling_price,
					g.tag`).
		Joins("inner join mall.tb_newbee_mall_goods_info g on c.goods_id = g.goods_id").
		Where("c.config_type = ?", typed).
		Find(&goods).Error
	if err != nil {
		return nil, xerrors.Wrapf(err, "")
	}
	return goods, nil
}
func (this *Repo) FetchAllCategory(ctx context.Context) ([]model.MallCategory, error) {
	category := make([]model.MallCategory, 0)
	err := this.db.Model(&model.MallCategory{}).Find(&category).Error
	if err != nil {
		return nil, xerrors.Wrapf(err, "failed to fetchAllCategory")
	}
	return category, nil
}
