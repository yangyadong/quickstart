package model

import (
	"github.com/astaxie/beego/logs"
	"quickstart/common"
	"time"
)

type Lottery struct {
	Id        int       `gorm:"id" json:"id"`
	Phone     string    `gorm:"phone" json:"phone"`
	Prize     int       `gorm:"prize" json:"prize"`
	Status    int       `gorm:"num" json:"num"`
	CreatedAt time.Time `gorm:"created_at" json:"created_at"`
	UpdatedAt time.Time `gorm:"updated_at" json:"updated_at"`
}

func AddLottery(lotteryData Lottery) error {
	err := common.Db.Create(&lotteryData).Error
	if err != nil {
		return err
	}
	return nil
}

func GetLotteryInfo(phone string) Lottery {
	today := time.Now().Format("2006-01-02")
	lottery := Lottery{}
	err := common.Db.
		Where("phone = ?", phone).
		Where("created_at >= ?", today).
		Last(&lottery).Error
	if err != nil {
		logs.Info("get user is fail: %v", err)
	}
	return lottery
}

// 获取某个用户中奖次数
func GetPhoneLotteryCount(phone string, prize int) int {
	var lottery []Lottery
	var count int
	err := common.Db.
		Where("status = ?", 1).
		Where("phone = ?", phone).
		Where("prize = ?", prize).
		Find(&lottery).Count(&count).Error
	if err != nil {
		logs.Info("get phone prize is fail: %v", err)
	}
	return count
}

// 获取某个奖项被抽走了几次
func GetLotteryCount(prize int) int {
	var lottery []Lottery
	var count int
	err := common.Db.
		Where("status = ?", 1).
		Where("prize = ?", prize).
		Find(&lottery).Count(&count).Error
	if err != nil {
		logs.Info("get prize is fail: %v", err)
	}
	return count
}
