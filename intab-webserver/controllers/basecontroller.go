package controllers

import (
	"html/template"
	pp "intab-core/passport"
	"intab-webserver/bootstrap"
	"log"

	srv "intab-core/services"

	"github.com/gorilla/sessions"
	"github.com/imdario/mergo"
	"github.com/rushairer/ago"
)

//ViewData 传递到view的数据集合
type ViewData map[string]interface{}

//BaseController 控制器基类
type BaseController struct {
	ago.Controller
	Passport   *pp.Passport
	layoutName string
	viewData   ViewData
	Session    *sessions.Session
	Loginned   bool
}

//Prepare 运行在每个Action之前
func (c *BaseController) Prepare() {

	c.Session, _ = srv.Session().Get(c.Request, srv.SessionID)

	cdnhost := bootstrap.GetApp().Config.CDNHost

	c.Passport = pp.InitPassport()
	c.Loginned = c.Passport.Get(c.Request, c.ResponseWriter)
	if c.Loginned {
		//log.Println(c.Passport.AccountRepository.User)
		c.viewData = ViewData{
			"IT_CDN_HOST": cdnhost,
			"currentUser": c.Passport.AccountRepository.User,
		}
	} else {
		c.viewData = ViewData{
			"IT_CDN_HOST": cdnhost,
		}
	}
}

//SaveSession 保存对Session做的修改
func (c *BaseController) SaveSession() {
	c.Session.Save(c.Request, c.ResponseWriter)
}

func (c *BaseController) layout(layoutName string) {
	c.layoutName = layoutName
}

func (c *BaseController) renderView(viewname string, data ViewData) {
	var err error
	var t *template.Template

	if len(c.layoutName) > 0 {
		t, err = template.ParseFiles("views/layouts/"+c.layoutName, "views/"+viewname)
	} else {
		t, err = template.ParseFiles("views/" + viewname)
	}
	if err != nil {
		log.Println(err)
	}

	if err := mergo.Merge(&c.viewData, data); err != nil {
		log.Println(err)
	}

	t.Execute(c.ResponseWriter, c.viewData)
}
