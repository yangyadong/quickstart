package model

import (
	"github.com/astaxie/beego/logs"
	"quickstart/common"
	"time"
)

type User struct {
	Id int  `gorm:"id" json:"id"`
	Phone string `gorm:"phone" json:"phone"`
	Desc string	`gorm:"desc" json:"desc"`
	CreatedAt time.Time `gorm:"created_at" json:"created_at"`
	UpdatedAt time.Time `gorm:"updated_at" json:"updated_at"`
}

func AddUser(userData User) error {
	err := common.Db.Create(&userData).Error
	if err != nil {
		return err
	}
	return nil
}

func GetUser(phone string) User {
	user := User{}
	err := common.Db.
		Where("phone = ?", phone).
		Last(&user).Error
	if err != nil {
		logs.Info("get user is fail: %v", err)
	}
	return user
}

func GetUserList(page int, limit int) (int, []User) {
	offset := (page - 1) * limit
	var users []User
	var count int
	err := common.Db.Find(&users).Count(&count).Error
	if err != nil {
		logs.Info("get user is fail: %v", err)
		return count, users
	}
	err = common.Db.
		Offset(offset).
		Limit(limit).
		Find(&users).Error
	if err != nil {
		logs.Info("get user is fail: %v", err)
	}
	return count, users
}