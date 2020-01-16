package main

import (
	bootstrap "intab-apiserver/bootstrap"
	c "intab-apiserver/controllers"
)

//InitRoutes 初始化路由
func InitRoutes() {
	app := bootstrap.GetApp()

	app.Ago.AddRoute(&c.HomeController{}, "/")

	//Account
	//app.Ago.AddRoute(&c.AccountController{}, "/account", "POST", "Create")

	//Document
	//app.Ago.AddRoute(&c.DocumentController{}, "/document/sharekey", "POST", "CreateSharekey").Middleware(middleware.Auth)
	//app.Ago.AddRoute(&c.DocumentController{}, "/document/{documentID:[0-9]+}", "GET", "GetOne").Middleware(middleware.Auth)
	//app.Ago.AddRoute(&c.DocumentController{}, "/document", "POST", "Create").Middleware(middleware.Auth)

	//OAuth2
	app.Ago.AddRoute(&c.OAuth2Controller{}, "/oauth2/token", "GET", "Token")
	app.Ago.AddRoute(&c.OAuth2Controller{}, "/oauth2/token", "POST", "Token")
}
