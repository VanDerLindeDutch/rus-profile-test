package main

import (
	"rus-profile-test/internal/app/grpc"
	"rus-profile-test/internal/app/rest"
	"rus-profile-test/internal/config"
)

func main() {
	cfg := config.GetConfig()
	mainApp := rest_app.NewApp(cfg)
	grpcApp := grpc_app.NewApp(cfg)
	go func() {
		err := grpcApp.Run()
		if err != nil {
			panic(err)
		}
	}()
	err := mainApp.Run()
	if err != nil {
		panic(err)
	}

}
