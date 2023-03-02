package utils

import (
	"math/rand"
	"time"
)

var (
	codeStr = []rune("1234567890")
)

func GetRandNum(n int) string {
	data := make([]rune, n)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := range data {
		data[i] = codeStr[r.Intn(len(codeStr))]
	}
	return string(data)
}
