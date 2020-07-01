package model

import (
	"encoding/json"
	"github.com/astaxie/beego/logs"
	"quickstart/common"
	"time"
)

type Prize struct {
	Id             int       `gorm:"id" json:"id"`
	Name           string    `gorm:"name" json:"name"`
	Num            int       `gorm:"num" json:"num"`
	Probability    int       `gorm:"probability" json:"probability"`
	NumOfUser      int       `gorm:"num_of_user" json:"num_of_user"`
	NumOfDay       int       `gorm:"num_of_day" json:"num_of_day"`
	IssuedQuantity int       `gorm:"Issued_quantity" json:"Issued_quantity"`
	Status         int       `gorm:"status" json:"status"`
	Desc           string    `gorm:"desc" json:"desc"`
	StartAt        time.Time `gorm:"start_at" json:"start_at"`
	EndAt          time.Time `gorm:"end_at" json:"end_at"`
	CreatedAt      time.Time `gorm:"created_at" json:"created_at"`
	UpdatedAt      time.Time `gorm:"updated_at" json:"updated_at"`
}

func GetPrizeList() []Prize {
	yesterday := time.Now().Add(1 * time.Hour)
	nowTime := time.Now()
	var prizes []Prize
	err := common.Db.
		Where("start_at <= ?", yesterday).
		Where("end_at >= ?", nowTime).
		Where("status = ?", 1).
		Find(&prizes).Error
	if err != nil {
		logs.Info("get user is fail: %v", err)
	}
	prizesJson, err := json.Marshal(prizes)
	if err != nil {
		logs.Error("prizes json marshal is err: %v", err)
	} else {
		common.RedisClient.Set("prize_cache", string(prizesJson), 5*time.Minute)
	}
	return prizes
}
