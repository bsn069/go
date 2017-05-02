package models

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/astaxie/beego"
)

type BaiduModel struct {
	strClientId     string
	strRedirectUri  string
	strClientSecret string
}

var G_BaiduModel BaiduModel

func (this *BaiduModel) UrlGetToken(strCode string) string {
	strAPIKey := beego.AppConfig.String("baidu_api_key")
	strRedirectUri := beego.AppConfig.String("baidu_redirect_uri")
	strSecretKey := beego.AppConfig.String("baidu_secret_key")

	strUrlFormat := `https://openapi.baidu.com/oauth/2.0/token?grant_type=authorization_code&code=%s&client_id=%s&client_secret=%s&redirect_uri=%s`
	strUrl := fmt.Sprintf(strUrlFormat, strCode, strAPIKey, strSecretKey, strRedirectUri)

	return strUrl
}

func (this *BaiduModel) UrlGetCode() string {
	strAPIKey := beego.AppConfig.String("baidu_api_key")
	strRedirectUri := beego.AppConfig.String("baidu_redirect_uri")

	strUrlFormat := `https://openapi.baidu.com/oauth/2.0/authorize?response_type=code&state=%s&client_id=%s&redirect_uri=%s&scope=basic`
	strUrl := fmt.Sprintf(strUrlFormat, "1", strAPIKey, strRedirectUri)

	return strUrl
}

type BaiduRsp struct {
	ExpiresIn     uint64 `json:"expires_in"`
	RefreshToken  string `json:"refresh_token"`
	AccessToken   string `json:"access_token"`
	Scope         string `json:"scope"`
	SessionSecret string `json:"session_secret"`
	SessionKey    string `json:"session_key"`
}

func NewBaiduRsp(strJson []byte) *BaiduRsp {
	pBaiduRsp := new(BaiduRsp)
	err := json.Unmarshal(strJson, pBaiduRsp)
	if err != nil {
		fmt.Println(err)
	}
	return pBaiduRsp
}

type BaiduUser struct {
	Id       string `json:"uid"`
	Name     string `json:"uname"`
	Portrait string `json:"portrait"`
}

func NewBaiduUser(strAccessToken string) *BaiduUser {
	strUrlFormat := `https://openapi.baidu.com/rest/2.0/passport/users/getLoggedInUser?access_token=%s&`
	strUrl := fmt.Sprintf(strUrlFormat, strAccessToken)
	resp, err := http.Get(strUrl)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
	pBaiduUser := new(BaiduUser)
	err = json.Unmarshal(body, pBaiduUser)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return pBaiduUser
}
