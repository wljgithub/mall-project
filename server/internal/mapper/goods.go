package mapper

import (
	"github.com/wljgithub/mall-project/internal/dto"
	"github.com/wljgithub/mall-project/internal/model"
)

func GoodsModelToGetGoodsDetailDto(goods model.Goods) dto.GetGoodsDetailRsp {
	return dto.GetGoodsDetailRsp{
		GoodsCoverImg:      goods.GoodsCoverImg,
		GoodsDetailContent: goods.GoodsDetailContent,
		GoodsId:            goods.GoodsId,
		GoodsIntro:         goods.GoodsIntro,
		GoodsName:          goods.GoodsName,
		OriginalPrice:      goods.OriginalPrice,
		SellingPrice:       goods.SellingPrice,
		Tag:                goods.Tag,
	}
}
