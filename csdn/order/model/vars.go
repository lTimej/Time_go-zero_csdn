package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var ErrNotFound = sqlx.ErrNotFound

// HomestayOrder 交易状态 :  0: 已取消 1:待支付 2:未使用 3:已使用  4:已过期
var OrderTradeStateCancel int64 = 0
var OrderTradeStateWaitPay int64 = 1
var OrderTradeStateWaitUse int64 = 2
var OrderTradeStateUsed int64 = 3
var OrderTradeStateExpire int64 = 4
