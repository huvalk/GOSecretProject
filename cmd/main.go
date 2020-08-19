package main

import (
	"GOSecretProject/core/server"
)

func main() {
	app := server.NewApp()
	app.StartRouter()
}