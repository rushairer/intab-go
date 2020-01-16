package controllers

import (
	pp "intab-core/passport"

	srv "intab-core/services"

	"github.com/gorilla/sessions"
	"github.com/rushairer/ago"
)

//ViewData 传递到view的数据集合
type ViewData map[string]interface{}

//BaseController 控制器基类
type BaseController struct {
	ago.Controller
	Passport *pp.Passport
	Session  *sessions.Session
}

//Prepare 运行在每个Action之前
func (c *BaseController) Prepare() {

	c.Session, _ = srv.Session().Get(c.Request, srv.SessionID)

	c.Passport = pp.InitPassport()
	c.Passport.Get(c.Request, c.ResponseWriter)

}

//SaveSession 保存对Session做的修改
func (c *BaseController) SaveSession() {
	c.Session.Save(c.Request, c.ResponseWriter)
}
