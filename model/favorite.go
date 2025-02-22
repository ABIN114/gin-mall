package model

import "gorm.io/gorm"

type Favorite struct {
	gorm.Model
	User      User    `gorm:"foreignKey:UserId"`
	UserId    uint    `gorm:"not null"`
	Product   Product `gorm:"foreignKey:ProductId"`
	ProductId uint    `gorm:"not null"`
	Boss      User    `gorm:"foreignKey:BossId"`
	BoosId    uint    `gorm:"not null"`
}
