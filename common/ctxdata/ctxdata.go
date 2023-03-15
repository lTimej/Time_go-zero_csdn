package ctxdata

import (
	"context"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"github.com/zeromicro/go-zero/core/logx"
	"strconv"
)

var CtxKeyJwtUserId = "jwtUserId"

type MyClaims struct {
	UserId string `json:"user_id"`
	jwt.RegisteredClaims
}

func GetUidFromCtx(ctx context.Context) int64 {
	user_id := ctx.Value(CtxKeyJwtUserId).(string)
	uid, err := strconv.ParseInt(user_id, 10, 64)
	if err != nil {
		logx.WithContext(ctx).Errorf("GetUidFromCtx err : %+v", err)
	}
	return uid
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
		fmt.Println(err, "虎虎虎虎虎虎")
		return nil, err
	}
	if claims, ok := token.Claims.(*MyClaims); ok && token.Valid {
		return claims, nil
	}
	fmt.Println("喜喜喜喜喜喜")
	return nil, errors.New("invalid token")
}
