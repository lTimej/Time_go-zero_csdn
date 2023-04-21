// Code generated by goctl. DO NOT EDIT.
package types

type Sku struct {
	SkuId  int64  `json:"sku_id"`
	SpecId string `json:"spec_id"`
	Specs  string `json:"specs"`
	Count  int64  `json:"count"`
}

type OrderCreateRequest struct {
	AddressId  int64   `json:"address_id"`
	TotalCount int64   `json:"total_count"`
	TotalPrice float32 `json:"total_price"`
	Sku        []Sku   `json:"sku"`
}

type OrderCreateResponse struct {
	Sn string `json:"sn"`
}

type OrderGetRequest struct {
}

type Spu struct {
	SkuId  int64   `json:"sku_id"`
	SpecId string  `json:"spec_id"`
	Specs  string  `json:"specs"`
	Count  int64   `json:"count"`
	Price  float32 `json:"price"`
}

type Orders struct {
	DefaultImage string  `json:"default_image"`
	TotalPrice   float32 `json:"total_price"`
	TotalCount   int64   `json:"total_count"`
	Sn           string  `json:"sn"`
	Freight      float32 `json:"freight"`
	OrderId      int64   `json:"order_id"`
	AddressId    int64   `json:"address_id"`
	PayStatus    int64   `json:"pay_status"`
	Spu          []Spu   `json:"spu"`
}

type OrderGetResponse struct {
	Orders []Orders `json:"orders"`
}
