package model

import (
	"time"
)

type Lottery struct {
	Id             int       `gorm:"id" json:"id"`
	Name           string    `gorm:"name" json:"name"`
	Num            int       `gorm:"num" json:"num"`
	IssuedQuantity int       `gorm:"Issued_quantity" json:"Issued_quantity"`
	Status         int       `gorm:"num" json:"num"`
	Desc           string    `gorm:"name" json:"name"`
	StartAt        time.Time `gorm:"created_at" json:"created_at"`
	EndAt          time.Time `gorm:"updated_at" json:"updated_at"`
	CreatedAt      time.Time `gorm:"created_at" json:"created_at"`
	UpdatedAt      time.Time `gorm:"updated_at" json:"updated_at"`
}
