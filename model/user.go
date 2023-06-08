package model

import "gorm.io/gorm"

// 用户表
type User struct {
	ID            int64  `gorm:"primaryKey"`
	Phone         string `gorm:"type:varchar(20);unique;not null;"`
	Password      string `gorm:"type:varchar(100);not null;default:''"`
	Name          string `gorm:"type:varchar(100)"`
	Avatar        string `gorm:"type:varchar(255)"`
	LastLoginTime int64  `gorm:"autoCreatedTime"`
	CreatedAt     int64  `gorm:"autoCreatedTime"`
	UpdatedAt     int64  `gorm:"autoCreatedTime"`
}

func IsRegister(db *gorm.DB, phone string) bool {
	var user User
	db.Where("phone = ?", phone).First(&user)
	return user.ID != 0
}
