package dto

type MallIndexGoods struct {
	GoodsCoverImg string `json:"goodsCoverImg"`
	GoodsId       int    `json:"goodsId"`
	GoodsInfo     string `json:"goodsInfo"`
	GoodsName     string `json:"goodsName"`
	SellingPrice  int    `json:"sellingPrice"`
	Tag           string `json:"tag"`
}
type CarouselItem struct {
	CarouselUrl string `json:"carouselUrl"`
	RedirectUrl string `json:"redirectUrl"`
}

type MallIndexRsp struct {
	Carousels        []CarouselItem   `json:"carousels"`
	HotGoodses       []MallIndexGoods `json:"hotGoodses"`
	NewGoodses       []MallIndexGoods `json:"newGoodses"`
	RecommendGoodses []MallIndexGoods `json:"recommendGoodses"`
}

type Categories struct {
	CategoryId    int          `json:"categoryId"`
	CategoryLevel int          `json:"categoryLevel"`
	CategoryName  string       `json:"categoryName"`
	SubCategories []Categories `json:"subCategories,omitempty"`
}
