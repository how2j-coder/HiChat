package main

import (
	"com/chat/service/cmd/chat/initial"
	"com/chat/service/pkg/app"
)

func main() {
	initial.InitApp()

	servers := initial.CreateService()
	shutdowns := initial.Shutdown(servers)
	a := app.New(servers, shutdowns)
	a.Run()
}
