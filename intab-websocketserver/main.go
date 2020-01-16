package main

import (
	bootstrap "intab-websocketserver/bootstrap"
)

func main() {
	app := bootstrap.GetApp()
	InitRoutes()
	app.Run()
}
