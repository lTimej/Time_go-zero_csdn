package ctxdata

import (
	"context"
	"errors"

	"fmt"

	"github.com/golang-jwt/jwt/v4"
	// "github.com/zeromicro/go-zero/core/logx"
)

var CtxKeyJwtUserId = "jwtUserId"

type MyClaims struct {
	UserId string `json:"user_id"`
	jwt.RegisteredClaims
}

func GetUidFromCtx(ctx context.Context) (user_id string) {
	if _, ok := ctx.Value(CtxKeyJwtUserId).(string); !ok {
		fmt.Println("获取userid失败")
		return ""
	}
	return ctx.Value(CtxKeyJwtUserId).(string)
}

//func GetUidFromCtx(ctx context.Context) int64 {
//	var uid int64
//	if jsonUid, ok := ctx.Value(CtxKeyJwtUserId).(json.Number); ok {
//		if int64Uid, err := jsonUid.Int64(); err == nil {
//			uid = int64Uid
//		} else {
//			logx.WithContext(ctx).Errorf("GetUidFromCtx err : %+v", err)
//		}
//	}
//	fmt.Println("%%%%%%^user_id%%%%%%%%%%%%", uid)
//	return uid
//}

func ParseToken(tokenString string, secretKey string) (*MyClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &MyClaims{}, func(token *jwt.Token) (i interface{}, err error) {
		return []byte(secretKey), nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*MyClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("invalid token")
}
