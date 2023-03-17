package utils

import (
	"strconv"
)

func StringToInt64(str string) int64 {
	if str == "" {
		return 0
	}
	res, _ := strconv.ParseInt(str, 10, 64)
	return res
}
