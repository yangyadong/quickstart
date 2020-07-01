package controllers

import (
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
	desc := c.GetString("desc")
	//captcha := c.GetString("captcha")
	user := model.User{Phone: phone, Desc: desc, CreatedAt: time.Now(), UpdatedAt: time.Now()}
	model.AddUser(user)
	c.Data["json"] = common.SendResponse(200, "success", "")
	c.ServeJSON()
}
