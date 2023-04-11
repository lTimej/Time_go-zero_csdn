// Code generated by goctl. DO NOT EDIT.
package types

type ProductSpuList struct {
	Name         string  `json:"name"`
	DefaultImage string  `json:"default_image"`
	Sales        int64   `json:"sales"`
	Cfavs        int64   `json:"cfavs"`
	SpuId        int64   `json:"spu_id"`
	Price        float32 `json:"price"`
	NowPrice     float32 `json:"now_price"`
}

type SubCategory struct {
	Name         string `json:"name"`
	Id           int64  `json:"id"`
	DefaultImage string `json:"default_image"`
}

type CategoryList struct {
	Name         string        `json:"name"`
	Id           int64         `json:"id"`
	SubCategorys []SubCategory `json:"subcategory"`
}

type ProductSpuListRequest struct {
	CategoryId int64 `json:"category_id,optional"`
}

type ProductSpuListResponse struct {
	ProductSpus []ProductSpuList `json:"product_spus"`
}

type ProductCategoryRequest struct {
}

type ProductCategoryResponse struct {
	Categorys []CategoryList `json:"categorys"`
}

type SpuDesc struct {
	DescInfo  string   `json:"desc_info"`
	DescImage []string `json:"desc_image"`
}

type Specs struct {
	SkuId     int64  `json:"sku_id"`
	SkuImg    string `json:"sku_img"`
	Name      string `json:"name"`
	SpecOptId int64  `json:"spec_opt_id"`
}

type SpecList struct {
	SpecId    int64   `json:"spec_id"`
	LabelName string  `json:"label_name"`
	Specs     []Specs `json:"specs"`
}

type SkuSpec struct {
	Label    string     `json:"label"`
	Stock    int64      `json:"stock"`
	SpecList []SpecList `json:"spec_list"`
}

type ProductDescRequest struct {
	SpuId int64 `json:"spu_id,optional"`
}

type ProductDescResponse struct {
	Title        string   `json:"title"`
	Price        float32  `json:"price"`
	NowPrice     float32  `json:"now_price"`
	Address      string   `json:"address"`
	SwiperImages []string `json:"swiper_images"`
	SpuDesc      SpuDesc  `json:"spu_desc"`
	SkuSpec      SkuSpec  `json:"sku_spec"`
}

type AddCartRequest struct {
	SkuId []int64 `json:"sku_id"`
	Count int64   `json:"count"`
}

type AddCartResponse struct {
}

type GetCartRequest struct {
}

type SpecLabel struct {
	Label string `json:"label"`
	Name  string `json:"name"`
}

type Carts struct {
	Title        string      `json:"title"`
	DefaultImage string      `json:"default_image"`
	Count        int64       `json:"count"`
	Price        float32     `json:"price"`
	SpecLabel    []SpecLabel `json:"spec_label"`
}

type GetCartResponse struct {
	Carts []Carts `json:"carts"`
}
