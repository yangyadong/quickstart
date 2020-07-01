package model

import (
	"quickstart/common"
	"time"
)

type User struct {
	id int  `gorm:"id" json:"id"`
	Phone string `gorm:"phone" json:"phone"`
	Desc string	`gorm:"desc" json:"desc"`
	CreatedAt time.Time `gorm:"created_at" json:"created_at"`
	UpdatedAt time.Time `gorm:"updated_at" json:"updated_at"`
}

func AddUser(userData User)  {
	common.Db.Create(&userData)
}