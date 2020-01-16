package main

import (
	bootstrap "intab-webserver/bootstrap"
)

func main() {
	app := bootstrap.GetApp()
	InitRoutes()
	app.Run()
}
