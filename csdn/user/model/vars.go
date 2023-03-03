package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var ErrNotFound = sqlx.ErrNotFound
var UserAuthTypePhone string = "phone"
var UserAuthTypeUsername string = "username"
