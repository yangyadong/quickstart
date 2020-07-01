package routers

import (
	"github.com/astaxie/beego"
	"quickstart/controllers"
)

func init() {
    //beego.Router("/", &controllers.MainController{})
	beego.Router("/participate", &controllers.ParticipateController{})
	beego.Router("/user", &controllers.UserController{})
	beego.Router("/captcha", &controllers.CaptchaController{})
	beego.Router("/lottery", &controllers.LotteryController{})
	beego.Router("/lottery/list", &controllers.LotteryListController{})
	beego.Router("/export/lottery", &controllers.ExportLotteryController{})
}
