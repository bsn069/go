package routers

import (
	"github.com/astaxie/beego"
	"github.com/bsn069/go/qq_login/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/qq", &controllers.QQController{})
	beego.Router("/baidu", &controllers.BaiduController{})
}
