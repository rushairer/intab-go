package main

import (
	agows "intab-websocketserver/agows"
	"intab-websocketserver/bootstrap"
	c "intab-websocketserver/controllers"
)

//InitRoutes 初始化路由
func InitRoutes() {
	app := bootstrap.GetApp()

	app.WebSocketServer.AddRoute(&c.DocumentController{}, "/document/get", agows.WebSocketCommandMsg, "GetDocuemtContent")
	app.WebSocketServer.AddRoute(&c.DocumentController{}, "/document/commit", agows.WebSocketCommandMsg, "Commit")
	app.WebSocketServer.AddRoute(&c.HistoryController{}, "/history/list", agows.WebSocketCommandMsg, "GetList")
}
