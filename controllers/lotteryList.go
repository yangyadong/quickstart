package controllers

import (
	"github.com/astaxie/beego"
	"quickstart/model"
)

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
