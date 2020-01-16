package main

import (
	bootstrap "intab-webserver/bootstrap"
	c "intab-webserver/controllers"
	middleware "intab-webserver/middlewares"
)

//InitRoutes 初始化路由
func InitRoutes() {
	app := bootstrap.GetApp()

	//Home
	app.Ago.AddRoute(&c.HomeController{}, "/")
	app.Ago.AddRoute(&c.HomeController{}, "/dashboard", "GET", "Dashboard").Middleware(middleware.Passport)
	//app.Ago.AddRoute(&c.HomeController{}, "/ws", "GET", "WS").Middleware(middleware.Passport)
	app.Ago.AddRoute(&c.HomeController{}, "/ws2", "GET", "WS2")

	//Document
	app.Ago.AddRoute(&c.DocumentController{}, "/document/new", "GET", "New").Middleware(middleware.Passport)
	app.Ago.AddRoute(&c.DocumentController{}, "/document/accessrequestlist", "GET", "AccessRequestList").Middleware(middleware.Passport)
	//访问单个文档，不需要登录验证
	app.Ago.AddRoute(&c.DocumentController{}, "/document/{key}", "GET", "GetOne")
	app.Ago.AddRoute(&c.DocumentController{}, "/document/{key}/request", "POST", "RequestAccess").Middleware(middleware.Passport)

	//Auth
	app.Ago.AddRoute(&c.AuthController{}, "/login", "GET", "Login")
	app.Ago.AddRoute(&c.AuthController{}, "/login", "POST", "PostLogin")
	app.Ago.AddRoute(&c.AuthController{}, "/logout", "GET", "Logout").Middleware(middleware.Passport)

	//OAuth2
	app.Ago.AddRoute(&c.OAuth2Controller{}, "/oauth2/wechat", "GET", "Wechat")
	app.Ago.AddRoute(&c.OAuth2Controller{}, "/oauth2/wechat_callback/{session_id}", "GET", "WechatCallback")

}
