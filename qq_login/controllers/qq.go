package controllers

import (
	"github.com/astaxie/beego"
)

type QQController struct {
	beego.Controller
}

func (c *QQController) Get() {
	c.Data["Website"] = "qq"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}
