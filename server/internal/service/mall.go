package service

import (
	"context"
	"github.com/wljgithub/mall-project/internal/dto"
	"github.com/wljgithub/mall-project/internal/mapper"
	"github.com/wljgithub/mall-project/internal/model"
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
