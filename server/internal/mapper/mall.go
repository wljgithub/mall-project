package mapper

import (
	"github.com/wljgithub/mall-project/internal/dto"
	"github.com/wljgithub/mall-project/internal/model"
)

func CarouselModelToCarousel(carousel []*model.Carousel) []dto.CarouselItem {
	dtoCarousel := make([]dto.CarouselItem, len(carousel))
	for index, item := range carousel {
		dtoCarousel[index].CarouselUrl = item.CarouselUrl
		dtoCarousel[index].RedirectUrl = item.RedirectUrl
	}
	return dtoCarousel
}
func MallCategoryToCategoryDto(categories []model.MallCategory) *dto.Categories {
	// 根据parentid 将数据归类
	mapper := make(map[int][]dto.Categories)
	for _, category := range categories {
		ele := dto.Categories{
			CategoryId:    category.CategoryId,
			CategoryLevel: category.CategoryLevel,
			CategoryName:  category.CategoryName,
		}
		mapper[category.ParentId] = append(mapper[category.ParentId], ele)
	}
	// 递归地构建多叉树
	var f func(*dto.Categories)
	f = func(root *dto.Categories) {
		root.SubCategories = mapper[root.CategoryId]
		for i := range root.SubCategories {
			f(&root.SubCategories[i])
		}
	}
	var root = &dto.Categories{}
	f(root)
	return root
}
