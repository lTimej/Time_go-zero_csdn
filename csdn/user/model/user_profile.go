package model

import (
	"time"
)

type UserProfile struct {
	UserId            int64     `gorm:"column:user_id"`             // 用户ID
	Gender            int64     `gorm:"column:gender"`              // 性别，0-男，1-女
	Birthday          time.Time `gorm:"column:birthday"`            // 生日
	RealName          string    `gorm:"column:real_name"`           // 真实姓名
	IdNumber          string    `gorm:"column:id_number"`           // 身份证号
	IdCardFront       string    `gorm:"column:id_card_front"`       // 身份证正面
	IdCardBack        string    `gorm:"column:id_card_back"`        // 身份证背面
	IdCardHandheld    string    `gorm:"column:id_card_handheld"`    // 手持身份证
	CreateTime        time.Time `gorm:"column:create_time"`         // 创建时间
	UpdateTime        time.Time `gorm:"column:update_time"`         // 更新时间
	RegisterMediaTime time.Time `gorm:"column:register_media_time"` // 注册自媒体时间
	Area              string    `gorm:"column:area"`                // 地区
	Company           string    `gorm:"column:company"`             // 公司
	Career            string    `gorm:"column:career"`              // 职业
	Tag               string    `gorm:"column:tag"`                 // 标签
}

func (t *UserProfile) TableName() string {
	return "user_profile"
}
