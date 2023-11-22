package rest_app

import (
	"context"
	"fmt"
	swagger_ui "github.com/esurdam/go-swagger-ui"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net/http"
	"os"
	"rus-profile-test/internal/config"
	"rus-profile-test/pkg/profile_v1"
)

type App struct {
	log *log.Logger
	cfg *config.Config
}

func NewApp(cfg *config.Config) *App {
	return &App{
		log: log.Default(),
		cfg: cfg,
	}
}

func (a *App) Run() error {

	ctx := context.Background()

	dopts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}

	swaggerFile, err := os.ReadFile(a.cfg.Swagger.FilePath)
	if err != nil {
		return err
	}
	mux := swagger_ui.NewServeMux(swagger_ui.AssetFnFromBytes(swaggerFile), "swagger.json")

	gwmux := runtime.NewServeMux()

	err = profile_v1.RegisterProfilerHandlerFromEndpoint(ctx, gwmux, fmt.Sprintf("localhost:%d", a.cfg.GRPC.Port), dopts)
	if err != nil {
		return err
	}

	mux.Handle("/", gwmux)

	fmt.Printf("REST server is running on port: %d\n", a.cfg.REST.Port)
	err = http.ListenAndServe(fmt.Sprintf(":%d", a.cfg.REST.Port), mux)
	if err != nil {
		return err
	}

	return nil
}
