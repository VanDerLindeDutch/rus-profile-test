package grpc_app

import (
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	"rus-profile-test/internal/config"
	"rus-profile-test/internal/domain/rus_profile"
	"rus-profile-test/internal/grpc/profiler"
)

type App struct {
	log    *log.Logger
	cfg    *config.Config
	server *grpc.Server
}

func NewApp(cfg *config.Config) *App {
	server := grpc.NewServer()

	grpc_profiler.Register(server, rus_profile.NewService(cfg))
	return &App{
		log:    log.Default(),
		cfg:    cfg,
		server: server,
	}
}

func (a *App) Run() error {
	l, err := net.Listen("tcp", fmt.Sprintf(":%d", a.cfg.GRPC.Port))
	if err != nil {
		return err
	}
	a.log.Println(fmt.Sprintf("Grpc server is running on port %d", a.cfg.GRPC.Port))
	if err := a.server.Serve(l); err != nil {
		return err
	}
	return nil
}
