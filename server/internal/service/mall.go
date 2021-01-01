package service

import (
	"context"
	"errors"
	"github.com/wljgithub/mall-project/internal/dto"
	"github.com/wljgithub/mall-project/internal/mapper"
	"github.com/wljgithub/mall-project/internal/model"
	"github.com/wljgithub/mall-project/internal/repository"
	"github.com/wljgithub/mall-project/pkg/errno"
	"golang.org/x/sync/errgroup"
	"time"
)

const (
	HotGoods = 3 + iota
	NewGoods
	RecommendGoods
)

func (this *Service) GetMallIndex() (*dto.MallIndexRsp, error) {
	mallIndex := &dto.MallIndexRsp{}
	timeoutCtx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	g, _ := errgroup.WithContext(timeoutCtx)

	var (
		err                                error
		carousel                           []*model.Carousel
		hotGoods, newGoods, recommendGoods []dto.MallIndexGoods
	)
	g.Go(func() error {
		carousel, err = this.Repo.FetchAllUser(context.Background())
		return err
	})
	g.Go(func() error {
		hotGoods, err = this.Repo.FetchGoodsByType(context.Background(), HotGoods)
		return err
	})
	g.Go(func() error {
		newGoods, err = this.Repo.FetchGoodsByType(context.Background(), NewGoods)
		return err
	})
	g.Go(func() error {
		recommendGoods, err = this.Repo.FetchGoodsByType(context.Background(), RecommendGoods)
		return err
	})
	g.Wait()
	mallIndex.Carousels = mapper.CarouselModelToCarousel(carousel)
	mallIndex.HotGoodses = hotGoods
	mallIndex.NewGoodses = newGoods
	mallIndex.RecommendGoodses = recommendGoods
	return mallIndex, nil
}

func (this *Service) GetCategory() (*dto.Categories, error) {
	category, err := this.Repo.FetchAllCategory(context.Background())
	if err != nil {
		return nil, err
	}
	return mapper.MallCategoryToCategoryDto(category), nil
}
func (this *Service) GetGoodsDetail(goodsId string) (*dto.GetGoodsDetailRsp, error) {
	goodsDetail, err := this.Repo.GetGoods(context.Background(), goodsId)
	if err != nil {
		if errors.Is(err, repository.ErrNotFound) {
			return nil, errno.ErrRecordNotFound
		}
		return nil, err
	}
	goodsDto := mapper.GoodsModelToGetGoodsDetailDto(goodsDetail)
	return &goodsDto, nil
}
func (this *Service) GoodsSearch(req dto.GoodsSearchReq) (*dto.GoodsSearchRsp, error) {
	if req.PageNumber <= 0 {
		req.PageNumber = 1
	}
	if req.PageSize <= 0 {
		req.PageSize = 5
	}

	var offset, limit int
	offset = (req.PageNumber - 1) * req.PageSize
	limit = req.PageSize

	var (
		goods = &dto.GoodsSearchRsp{}
		err   error
	)
	if req.GoodsCategoryId != "" {
		goods.List, err = this.Repo.SearchGoodsByCaregory(context.Background(), req.GoodsCategoryId, req.OrderBy, offset, limit)
		goods.TotalCount, err = this.Repo.CountGoodsByCategory(context.Background(), req.GoodsCategoryId)
	} else {
		goods.List, err = this.Repo.SearchGoodsByName(context.Background(), req.Keyword, req.OrderBy, offset, limit)
		goods.TotalCount, err = this.Repo.CountGoodsByName(context.Background(), req.Keyword)
	}

	if err != nil {
		return nil, err
	}
	return goods, nil
}
