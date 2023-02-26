package utils

import (
	"github.com/golang-jwt/jwt/v4"
	"liujun/Time_go-zero_csdn/csdn/user/constants"
	"time"
)

func GenToken(user_id int64, second time.Duration) (string, error) {
	claim := constants.TokenClaim{
		user_id,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(second)),
			Issuer:    "cloud_disk", // 签发人
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	return token.SignedString(constants.TokenKey)
}
