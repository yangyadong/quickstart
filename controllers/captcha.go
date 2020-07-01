package controllers

import (
	"github.com/astaxie/beego"
	"quickstart/common"
	"time"
)

type CaptchaController struct {
	beego.Controller
}

func (c *CaptchaController) Post() {
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

func setPhoneCaptcha(phone string) bool {
	captcha := common.GenValidateCode(6)
	captchaIndex := getCaptchaIndex(phone)
	res, err := common.RedisClient.SetNX(captchaIndex, captcha, 15 * time.Minute).Result()
	if err != nil || !res{
		return false
	}
	common.SendCaptchaMsg(phone, captcha)
	return true
}

func checkCaptcha(phone string, captcha string) bool {
	if captcha == "" {
		return false
	}
	captchaIndex := getCaptchaIndex(phone)
	captchaRedis := common.RedisClient.Get(captchaIndex).Val()
	if captchaRedis != captcha {
		return false
	}
	common.RedisClient.Del(captchaIndex)
	return true
}

func getCaptchaIndex(phone string) string {
	return "phone_captcha: " + phone
}