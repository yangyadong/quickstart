package common

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/go-sql-driver/mysql"
)

var Db *gorm.DB

func InitMysql() {
	//连接数据库
	arrInfo := fmt.Sprintf("%v:%v@tcp(%v)/%v?charset=utf8", ConfigInfo.Mysql.User, ConfigInfo.Mysql.Pwd,ConfigInfo.Mysql.Addr,ConfigInfo.Mysql.Name)
	Db, _ = gorm.Open(ConfigInfo.Mysql.Dialect, arrInfo)
	Db.SingularTable(true)
}
