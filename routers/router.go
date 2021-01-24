package routers

import (
	"BaGua/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
	beego.Router("/qigua", &controllers.QiGuaController{}, "get:QiGua")
}
