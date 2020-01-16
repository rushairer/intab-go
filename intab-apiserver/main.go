package main

import (
	bootstrap "intab-apiserver/bootstrap"
	cmd "intab-apiserver/commands"
	"os"
)

func main() {
	app := bootstrap.GetApp()

	InitRoutes()

	args := os.Args
	if len(args) > 1 {
		switch {
		case args[1] == "migrate" && len(args) == 2:
			cmd.MigrateDB()
		case len(args) > 1:
			fallthrough
		case args[1] == "help":
			cmd.ShowHelp()
		}
	} else {
		app.Run()
	}
}
