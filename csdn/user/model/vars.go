package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var ErrNotFound = sqlx.ErrNotFound
var UserAuthTypePhone string = "phone"       //平台内部
var UserAuthTypeUsername string = "username" //微信小程序
