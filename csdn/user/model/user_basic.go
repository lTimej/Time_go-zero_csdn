package model

import (
	"time"
)

type UserBasic struct {
	UserId         int64     `gorm:"column:user_id"`         // 用户ID
	Account        string    `gorm:"column:account"`         // 账号
	Email          string    `gorm:"column:email"`           // 邮箱
	Status         int64     `gorm:"column:status"`          // 状态，是否可用，0-不可用，1-可用
	Mobile         string    `gorm:"column:mobile"`          // 手机号
	Password       string    `gorm:"column:password"`        // 密码
	UserName       string    `gorm:"column:user_name"`       // 昵称
	ProfilePhoto   string    `gorm:"column:profile_photo"`   // 头像
	LastLogin      time.Time `gorm:"column:last_login"`      // 最后登录时间
	IsMedia        int64     `gorm:"column:is_media"`        // 是否是自媒体，0-不是，1-是
	IsVerified     int64     `gorm:"column:is_verified"`     // 是否实名认证，0-不是，1-是
	Introduction   string    `gorm:"column:introduction"`    // 简介
	Certificate    string    `gorm:"column:certificate"`     // 认证
	ArticleCount   int64     `gorm:"column:article_count"`   // 发文章数
	FollowingCount int64     `gorm:"column:following_count"` // 关注的人数
	FansCount      int64     `gorm:"column:fans_count"`      // 被关注的人数
	LikeCount      int64     `gorm:"column:like_count"`      // 累计点赞人数
	ReadCount      int64     `gorm:"column:read_count"`      // 累计阅读人数
	CodeYear       int64     `gorm:"column:code_year"`       // 码龄
}

func (t *UserBasic) TableName() string {
	return "user_basic"
}
