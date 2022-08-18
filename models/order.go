package models

import "time"

type Order struct {
	Id        uint `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time
	ProductId int     `json:"productId"`
	Product   Product `gorm:"foreignKey:ProductId"`
	UserId    int     `json:"userId"`
	User      User    `gorm:"foreignKey:UserId"`
}
