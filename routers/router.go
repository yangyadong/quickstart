package routers

import (
	"github.com/astaxie/beego"
	"quickstart/controllers"
)

func init() {
    //beego.Router("/", &controllers.MainController{})
	beego.Router("/user", &controllers.UserController{})
	beego.Router("/captcha", &controllers.CaptchaController{})
}
