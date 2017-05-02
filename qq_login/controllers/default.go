package controllers

import (
	"github.com/astaxie/beego"
	"github.com/bsn069/go/qq_login/models"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["BaiduUrl"] = models.G_BaiduModel.UrlGetCode()

	c.TplName = "all.tpl"
}
