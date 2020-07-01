package common

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

type User struct {
	id int  `gorm:"id" json:"id"`
	Name string	`gorm:"name" json:"name"`
	Phone string `gorm:"phone" json:"phone"`
	CreatedAt *time.Time `gorm:"created_at" json:"created_at"`
	UpdatedAt *time.Time `gorm:"updated_at" json:"updated_at"`
}

var Db *gorm.DB

func InitMysql() {
	//连接数据库
	arrInfo := fmt.Sprintf("%v:%v@tcp(%v)/%v?charset=utf8", ConfigInfo.Mysql.User, ConfigInfo.Mysql.Pwd,ConfigInfo.Mysql.Addr,ConfigInfo.Mysql.Name)
	Db, err := gorm.Open(ConfigInfo.Mysql.Dialect, arrInfo)
	//一个坑，不设置这个参数，gorm会把表名转义后加个s，导致找不到数据库的表
	Db.SingularTable(true)
	if err != nil {
		panic(err)
	}
}
