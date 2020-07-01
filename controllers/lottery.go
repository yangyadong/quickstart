package controllers

import (
	"bytes"
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/tealeg/xlsx"
	"math/rand"
	"net/http"
	"quickstart/common"
	"quickstart/model"
	"strconv"
	"time"
)

type ExportLotteryController struct {
	beego.Controller
}

func (c *ExportLotteryController) Get() {
	file := xlsx.NewFile()
	sheet, _ := file.AddSheet("sheet")
	//设置表格头
	row := sheet.AddRow()
	var headers = []string{"序号", "账号", "奖品", "时间"}
	for _, header := range headers {
		row.AddCell().Value = header
	}
	_, dataArr := model.GetLotteryList(1, 10000)
	//写入数据
	for i, data := range dataArr {
		row := sheet.AddRow()
		row.AddCell().Value = strconv.Itoa(i + 1)
		row.AddCell().Value = data.Phone
		row.AddCell().Value = data.PrizeName
		row.AddCell().Value = data.CreatedAt.Format("2006-01-02 15:04:05")
	}
	c.Ctx.ResponseWriter.Header().Add("Content-Disposition", "attachment")
	c.Ctx.ResponseWriter.Header().Add("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")

	var buffer bytes.Buffer
	if err := file.Write(&buffer); err != nil {
		//return err
	}
	r := bytes.NewReader(buffer.Bytes())
	http.ServeContent(c.Ctx.ResponseWriter, c.Ctx.Request, "", time.Now(), r)
}

type LotteryListController struct {
	beego.Controller
}

func (c *LotteryListController) Get() {
	page, err := c.GetInt("page", 1)
	if err != nil || page < 1 {
		page = 1
	}
	limit, err := c.GetInt("limit", 10)
	if err != nil || limit < 1 {
		limit = 10
	}
	count, lotterys := model.GetLotteryList(page, limit)
	c.Data["lotterys"] = lotterys
	c.Data["count"] = count
	c.TplName = "lottery/lotteryList.tpl"
}

type LotteryController struct {
	beego.Controller
}

func (c *LotteryController) Get() {
	c.TplName = "lottery/lottery.tpl"
}

func (c *LotteryController) Post() {
	c.TplName = "lottery/lotteryRes.tpl"
	phone := c.GetString("phone")
	checkRes := common.CheckPhone(phone)
	if !checkRes {
		c.Data["json"] = common.SendResponse(400, "error", "phone number is illegal")
		//c.ServeJSON()
		c.Data["lotteryRes"] = "手机号格式有误"
		return
	}
	checkRes, _ = checkUser(phone)
	if checkRes {
		c.Data["json"] = common.SendResponse(400, "error", "this user is not registered")
		//c.ServeJSON()
		c.Data["lotteryRes"] = "请先报名参与"
		return
	}
	lotteryIndex := "phone_lottery: " + phone
	setRes, err := common.RedisClient.SetNX(lotteryIndex, "true", 5*time.Minute).Result()
	if err != nil || !setRes {
		c.Data["json"] = common.SendResponse(400, "error", "phone lottery fail")
		//c.ServeJSON()
		c.Data["lotteryRes"] = "错误，请稍后再试"
		return
	}
	defer common.RedisClient.Del(lotteryIndex)
	// 判断用户今天是否已经抽过奖了
	lottery := model.GetLotteryInfo(phone)
	if lottery.Id != 0 {
		c.Data["json"] = common.SendResponse(400, "error", "participated today")
		//c.ServeJSON()
		c.Data["lotteryRes"] = "今日已抽奖，请明天再来"
		return
	}
	isPrice, lotteryRes, priceId, priceName := getPrice(phone)
	if !isPrice {
		c.Data["json"] = common.SendResponse(500, "error", "Lottery failed")
		//c.ServeJSON()
		c.Data["lotteryRes"] = "活动未开始"
		return
	}
	lotteryData := model.Lottery{
		Status:    lotteryRes,
		Phone:     phone,
		Prize:     priceId,
		PrizeName: priceName,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	res := model.AddLottery(lotteryData)
	if res != nil {
		c.Data["json"] = common.SendResponse(400, "error", res)
		c.Data["lotteryRes"] = "失败，系统错误"
	} else {
		c.Data["json"] = common.SendResponse(200, "success", "")
		if lotteryRes == 1 {
			c.Data["lotteryRes"] = "恭喜获得 " + priceName + " 一份"
		} else {
			c.Data["lotteryRes"] = "再接再厉"
		}
	}
	//c.ServeJSON()
}

// 开始抽奖
func getPrice(phone string) (bool, int, int, string) {
	prizesJson, err := common.RedisClient.Get("prize_cache").Result()
	var prizes []model.Prize
	if err != nil || prizesJson == "" {
		prizes = model.GetPrizeList()
	} else {
		err := json.Unmarshal([]byte(prizesJson), &prizes)
		if err != nil {
			prizes = model.GetPrizeList()
		}
	}
	nowTime := time.Now()
	today := nowTime.Format("2006-01-02")
	var prizesEffective []model.Prize
	for _, prize := range prizes {
		if !(prize.StartAt.Unix() <= nowTime.Unix() && prize.EndAt.Unix() >= nowTime.Unix()) {
			continue
		}
		prizesEffective = append(prizesEffective, prize)
	}

	// 随机在一个奖项中抽奖
	// 默认未中奖
	lotteryRes := 0
	prizesEffectiveCount := len(prizesEffective)
	if prizesEffectiveCount == 0 {
		return false, lotteryRes, 0, ""
	}
	index := rand.Intn(prizesEffectiveCount)
	prize := prizesEffective[index]

	// 开始计算抽奖资格
	prizeIndex := today + "_" + strconv.Itoa(prize.Id)
	prizeNumOfDay, _ := common.RedisClient.Get(prizeIndex).Result()
	prizeNumOfDayCache, _ := strconv.Atoi(prizeNumOfDay)
	// 判断这个奖项今天是否还能抽奖【奖品没有每天上限个数，或还没达到上限个数】
	if prize.NumOfDay == -1 || (prizeNumOfDay == "" || prizeNumOfDayCache < prize.NumOfDay) {
		rand := rand.Intn(100) + 1
		if rand <= prize.Probability {
			lotteryRes = 1
			// 判断这个奖项被指定用户抽走了几次了，是否达到该奖项单个用户可获取上限
			if prize.NumOfUser != -1 {
				count := model.GetPhoneLotteryCount(phone, prize.Id)
				if count >= prize.NumOfUser {
					lotteryRes = 0
				}
			}
			// 判断这个奖项被抽走了几次了，是否达到该奖项可获取上限
			if lotteryRes == 1 && prize.Num != -1 {
				prizeId := "prize" + "_" + strconv.Itoa(prize.Id)
				prizeNumAll, err := common.RedisClient.Get(prizeId).Result()
				prizeNumAllCache, err1 := strconv.Atoi(prizeNumAll)
				if err != nil || err1 != nil || prizeNumAll == "" || prizeNumAllCache == 0 {
					prizeNumAllCache = model.GetLotteryCount(prize.Id)
				}
				if prizeNumAllCache >= prize.Num {
					lotteryRes = 0
				} else {
					prizeNumAllCache++
					common.RedisClient.Set(prizeId, prizeNumAllCache, 1*time.Hour)
				}
			}
			// 中奖了--将有每日获取上限的奖品，今日获取的总次数缓存到redis
			if prize.NumOfDay != -1 && lotteryRes == 1 {
				prizeNumOfDayCache += 1
				common.RedisClient.Set(prizeIndex, prizeNumOfDayCache, 24*time.Hour)
			}
		}
	}
	return true, lotteryRes, prize.Id, prize.Name
}
