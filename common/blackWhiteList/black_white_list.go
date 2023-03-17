package blackWhiteList

var BlackWhiteList []map[string]bool

func init() {
	BlackWhiteList = []map[string]bool{
		map[string]bool{"/v1/article/status": true},
	}
}
