package utils

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

func Md5ByString(str string) string {
	m := md5.New()
	_, err := m.Write([]byte(str))
	if err != nil {
		panic(err)
	}
	arr := m.Sum(nil)
	return hex.EncodeToString(arr)
}

func Md5ByBytes(b []byte) string {
	return fmt.Sprintf("%x", md5.Sum(b))
}
