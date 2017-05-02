package controllers

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/astaxie/beego"
	"github.com/bsn069/go/qq_login/models"
)

type BaiduController struct {
	beego.Controller
}

func (c *BaiduController) Get() {
	strCode := c.GetString("code")
	strState := c.GetString("state")
	fmt.Print(strState)
	c.TplName = "baidu.tpl"

	strUrl := models.G_BaiduModel.UrlGetToken(strCode)
	resp, err := http.Get(strUrl)
	if err != nil {
		fmt.Println(err)
		c.Data["Url"] = err
		return
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
	pBaiduRsp := models.NewBaiduRsp(body)
	fmt.Println(pBaiduRsp.AccessToken)

	pBaiduUser := models.NewBaiduUser(pBaiduRsp.AccessToken)
	if pBaiduUser == nil {
		fmt.Println("pBaiduUser == nil")
		c.Data["Url"] = "nil"
		return
	}

	c.Data["Url"] = pBaiduUser.Id
}
