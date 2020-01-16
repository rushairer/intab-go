package controllers

import (
	repo "intab-core/repositories"
)

//HomeController Home页面控制器
type HomeController struct {
	BaseController
}

//Get 首页
func (c *HomeController) Get() {
	c.layout("single.html")
	c.renderView("index.html", nil)
}

//Dashboard 控制面板
func (c *HomeController) Dashboard() {
	user := c.Passport.AccountRepository.User
	resourceRepository := repo.InitResourceRepository()

	//TODO: 暂时没有分页处理，日后优化
	resources := resourceRepository.ResourceListWithUserIDAndLimit(user.ID, 0, 1000)

	c.layout("master.html")
	c.renderView("dashboard.html", ViewData{
		"user":      user,
		"resources": resources,
		"pageTitle": "我的桌面",
	})
}

//WS WebSocket测试页
func (c *HomeController) WS() {
	user := c.Passport.AccountRepository.User
	tokenString := c.Passport.GetAccessTokenString()

	ch := c.Request.FormValue("ch")

	c.layout("master.html")
	c.renderView("ws.html", ViewData{
		"user":  user,
		"token": tokenString,
		"ch":    ch,
	})
}

//WS2 WebSocket测试页2
func (c *HomeController) WS2() {
	user := c.Passport.AccountRepository.User
	tokenString := c.Passport.GetAccessTokenString()

	ch := c.Request.FormValue("ch")

	c.layout("master.html")
	c.renderView("ws2.html", ViewData{
		"user":  user,
		"token": tokenString,
		"ch":    ch,
	})
}

//WS3 WebSocket测试页2
func (c *HomeController) WS3() {
	user := c.Passport.AccountRepository.User
	tokenString := c.Passport.GetAccessTokenString()

	ch := c.Request.FormValue("ch")

	c.layout("master.html")
	c.renderView("ws3.html", ViewData{
		"user":  user,
		"token": tokenString,
		"ch":    ch,
	})
}
