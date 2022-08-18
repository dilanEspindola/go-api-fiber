package models

import "time"

type User struct {
	Id        uint   `json:"id" gorm:"primaryKey"`
	Name      string `json:"Name"`
	LastName  string `json:"lastName"`
	CreatedAt time.Time
}
