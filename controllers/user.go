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
	c.Data["Website"] = "user"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}

func (c *UserController) Post() {
	phone := c.GetString("phone")
	checkRes := common.CheckPhone(phone)
	if !checkRes {
		c.Data["json"] = common.SendResponse(400, "error", "phone number is illegal")
		c.ServeJSON()
		return
	}
	captcha := c.GetString("captcha")
	checkRes = checkCaptcha(phone, captcha)
	if !checkRes {
		c.Data["json"] = common.SendResponse(400, "error", "captcha is fail")
		c.ServeJSON()
		return
	}
	checkRes, msg := checkUser(phone)
	if !checkRes {
		c.Data["json"] = common.SendResponse(400, "error", msg)
		c.ServeJSON()
		return
	}
	desc := c.GetString("desc")
	user := model.User{Phone: phone, Desc: desc, CreatedAt: time.Now(), UpdatedAt: time.Now()}
	res := model.AddUser(user)
	if res != nil {
		c.Data["json"] = common.SendResponse(400, "error", res)
	} else {
		c.Data["json"] = common.SendResponse(200, "success", "")
	}
	c.ServeJSON()
}

func checkUser(phone string) (bool, string){
	user := model.GetUser(phone)
	fmt.Println(user)
	if user.Id != 0 {
		return false, "phone number already exists"
	}
	return true, ""
}