package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"quickstart/common"
	"quickstart/model"
	"time"
)

type UserController struct {
	beego.Controller
}

func (c *UserController) Get() {
	page, err := c.GetInt("page", 1)
	if err != nil || page < 1 {
		page = 1
	}
	limit, err := c.GetInt("limit", 10)
	if err != nil || limit < 1 {
		limit = 10
	}
	count, users := model.GetUserList(page, limit)
	c.Data["users"] = users
	c.Data["count"] = count
	c.TplName = "user/userList.tpl"
}

func (c *UserController) Post() {
	c.TplName = "user/participateRes.tpl"
	phone := c.GetString("phone", "")
	checkRes := common.CheckPhone(phone)
	if !checkRes {
		c.Data["json"] = common.SendResponse(400, "error", "phone number is illegal")
		//c.ServeJSON()
		c.Data["participateRes"] = "手机号格式错误"
		return
	}
	captcha := c.GetString("captcha", "")
	checkRes = checkCaptcha(phone, captcha)
	if !checkRes {
		c.Data["json"] = common.SendResponse(400, "error", "captcha is fail")
		//c.ServeJSON()
		c.Data["participateRes"] = "手机验证码错误"
		return
	}
	checkRes, msg := checkUser(phone)
	if !checkRes {
		c.Data["json"] = common.SendResponse(400, "error", msg)
		//c.ServeJSON()
		c.Data["participateRes"] = "已报名"
		return
	}
	desc := c.GetString("desc", "")
	if desc == "" {
		c.Data["json"] = common.SendResponse(400, "error", "essay cannot be empty")
		//c.ServeJSON()
		return
	}
	user := model.User{Phone: phone, Desc: desc, CreatedAt: time.Now(), UpdatedAt: time.Now()}
	res := model.AddUser(user)
	if res != nil {
		c.Data["json"] = common.SendResponse(400, "error", res)
		c.Data["participateRes"] = "失败，系统故障"
	} else {
		c.Data["json"] = common.SendResponse(200, "success", "")
		c.Data["participateRes"] = "报名成功"
	}
	//c.ServeJSON()
}

func checkUser(phone string) (bool, string){
	user := model.GetUser(phone)
	fmt.Println(user)
	if user.Id != 0 {
		return false, "phone number already exists"
	}
	return true, ""
}