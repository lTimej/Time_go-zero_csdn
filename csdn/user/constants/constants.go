package constants

import "github.com/golang-jwt/jwt/v4"

const CodeExpire = 300

var TokenKey = []byte("cloud_disk")

type TokenClaim struct {
	UserId int64
	jwt.RegisteredClaims
}
