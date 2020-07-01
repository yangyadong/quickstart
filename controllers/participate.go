package controllers

import (
	"github.com/astaxie/beego"
)

type ParticipateController struct {
	beego.Controller
}

func (c *ParticipateController) Get() {
	c.TplName = "user/addUser.tpl"
}
