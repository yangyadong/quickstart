package controllers

import (
	"github.com/astaxie/beego"
	"quickstart/common"
)

type LotteryController struct {
	beego.Controller
}

func (c *LotteryController) Post() {
	phone := c.GetString("phone")
	checkRes := common.CheckPhone(phone)
	if !checkRes {
		c.Data["json"] = common.SendResponse(400, "error", "phone number is illegal")
		c.ServeJSON()
		return
	}
	res := setPhoneCaptcha(phone)
	if !res {
		c.Data["json"] = common.SendResponse(400, "error", "captcha is exist")
	} else {
		c.Data["json"] = common.SendResponse(200, "success", "")
	}
	c.ServeJSON()
}
