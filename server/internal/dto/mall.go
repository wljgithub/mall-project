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

// GET: /api/v1/goods/detail/{goodsId}
type GetGoodsDetailRsp struct {
	GoodsCoverImg      string `json:"goodsCoverImg"`
	GoodsDetailContent string `json:"goodsDetailContent"`
	GoodsId            int    `json:"goodsId"`
	GoodsIntro         string `json:"goodsIntro"`
	GoodsName          string `json:"goodsName"`
	OriginalPrice      int    `json:"originalPrice"`
	SellingPrice       int    `json:"sellingPrice"`
	Tag                string `json:"tag"`
}

// GET: /api/v1/search
type GoodsSearchReq struct {
	Keyword         string `form:"keyword"`
	GoodsCategoryId string `form:"goodsCategoryId"`
	OrderBy         string `form:"orderBy"`
	PageNumber      int    `form:"pageNumber"`
	PageSize        int    `form:"pageSize"`
}

type SearchGoodsItem struct {
	GoodsCoverImg string `json:"goodsCoverImg"`
	GoodsId int `json:"goodsId"`
	GoodsIntro string `json:"goodsIntro"`
	GoodsName string  `json:"goodsName"`
	SellingPrice int `json:"sellingPrice"`
}
type GoodsSearchRsp struct {
	List       []SearchGoodsItem `json:"list"`
	TotalCount int            `json:"totalCount"`
	TotalPage int `json:"totalPage"`
}
