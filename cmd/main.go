package main

import (
	"GOSecretProject/core/server"
	"github.com/kataras/golog"
)

func main() {
	app := server.NewApp()

	if app != nil {
		app.StartRouter()
	} else {
		golog.Fatal("start fatal")
	}
}
