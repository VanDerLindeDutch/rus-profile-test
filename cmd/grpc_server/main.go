package main

import (
	app "rus-profile-test/internal/app/grpc"
	"rus-profile-test/internal/config"
)

func main() {
	cfg := config.GetConfig()
	mainApp := app.NewApp(cfg)
	err := mainApp.Run()
	if err != nil {
		panic(err)
	}
}
