package models

import "time"

type Product struct {
	Id        uint   `json:"id" gorm:"primaryKey"`
	Name      string `json:"name"`
	Price     int    `json:"description"`
	CreatedAt time.Time
}
