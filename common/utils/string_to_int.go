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

func Int64ToString(n int64) string {
	if n == 0 {
		return "0"
	}
	return strconv.FormatInt(n, 10)
}
