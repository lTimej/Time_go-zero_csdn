package utils

import (
	"math/rand"
	"strings"
	"time"
)

var (
	codeStr   = []rune("1234567890")
	charModel = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
)

func GetRandNum(n int) string {
	data := make([]rune, n)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := range data {
		data[i] = codeStr[r.Intn(len(codeStr))]
	}
	return string(data)
}

func GetRandstring(length int) string {
	if length < 1 {
		return ""
	}
	charArr := strings.Split(charModel, "")
	charlen := len(charArr)
	ran := rand.New(rand.NewSource(time.Now().Unix()))

	var rchar string = ""
	for i := 1; i <= length; i++ {
		rchar = rchar + charArr[ran.Intn(charlen)]
	}
	return rchar
}
