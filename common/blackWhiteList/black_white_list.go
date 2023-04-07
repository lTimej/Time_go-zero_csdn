package blackWhiteList

var BlackWhiteList map[string]bool

func init() {
	BlackWhiteList = map[string]bool{
		"/v1/article/status": true,
		"/v1/shop/add/cart":  true,
	}
}
