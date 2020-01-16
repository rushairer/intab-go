package controllers

import (
	"crypto/md5"
	"fmt"
	"html"
	pp "intab-core/passport"
	repo "intab-core/repositories"
	srv "intab-core/services"
	b "intab-webserver/bootstrap"
	"net/http"

	"github.com/bitly/go-simplejson"
	"github.com/segmentio/ksuid"
)

//OAuth2Controller OAuth2控制器
type OAuth2Controller struct {
	BaseController
}

//Wechat oauth2/wechat 微信登录入口
func (c *OAuth2Controller) Wechat() {
	sessionID := ksuid.New().String()

	session, _ := srv.Session().Get(c.Request, sessionID)
	session.Values["value"] = 1
	session.Save(c.Request, c.ResponseWriter)

	callback := b.GetApp().Config.OAuth2.WechatCallbackURL
	callback = fmt.Sprintf(callback, c.Request.Host)
	callback = html.EscapeString(callback + "/" + sessionID)

	url := b.GetApp().Config.OAuth2.WechatAPIUrl
	url = url + "?callback=" + callback
	http.Redirect(c.ResponseWriter, c.Request, url, 302)

}

//WechatCallback oauth2/wechatcallback 微信登录入口回调
func (c *OAuth2Controller) WechatCallback() {
	sessionID := c.Params["session_id"]
	session, _ := srv.Session().Get(c.Request, sessionID)

	if session.Values["value"] == 1 {
		openID := c.Request.FormValue("open_id")

		session.Options.MaxAge = -1
		session.Save(c.Request, c.ResponseWriter)

		url := b.GetApp().Config.OAuth2.WechatUserInfoURL
		url = url + "/" + openID

		client := new(http.Client)
		req, _ := http.NewRequest("GET", url, nil)
		resp, err := client.Do(req)
		if err != nil {
			fmt.Println(err)
		}
		defer resp.Body.Close()

		json, _ := simplejson.NewFromReader(resp.Body)

		nickname := json.Get("nickname").MustString()
		avatar := json.Get("headimgurl").MustString()
		unionID := json.Get("unionID").MustString()
		//subscribe := json.Get("subscribe").MustInt()

		username := "wx_" + openID
		password := md5.Sum([]byte(unionID))

		accountRepository, err := repo.NewAccountRepository(username, "", "", fmt.Sprintf("%x", password))
		accountRepository.User.UserDetail.Nickname = nickname
		accountRepository.User.UserDetail.Avatar = avatar
		accountRepository.SetUserStatusNormal()

		if err == nil {
			//TODO: 优化用户扩展信息的同步刷新
			accountRepository.LoginOrCreate()
			passport := pp.Passport{}
			passport.AccountRepository = accountRepository
			passport.Store(c.Request, c.ResponseWriter)

			http.Redirect(c.ResponseWriter, c.Request, "/dashboard", 302)
		} else {
			c.HTML(err.Error())
		}

	} else {
		c.HTMLForbiddon()
	}
}
