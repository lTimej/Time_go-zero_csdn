package utils

import (
	"errors"
	"fmt"
	"liujun/Time_go-zero_csdn/common/ctxdata"

	"github.com/golang-jwt/jwt/v4"
)

func VerifyToken(token_str, secret_key string) (int64, error) {
	token, err := jwt.Parse(token_str, func(token *jwt.Token) (interface{}, error) {
		fmt.Println("@@@@@@@@@@@@@")
		return []byte(secret_key), nil
	})
	fmt.Println(token, "###############")
	if err != nil {
		return 0, err
	}
	if claim, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		user_id := int64(claim[ctxdata.CtxKeyJwtUserId].(float64))
		return user_id, nil
	}
	return 0, errors.New("非法token")
}
