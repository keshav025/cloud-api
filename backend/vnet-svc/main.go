package main

import (
	"context"
	"net/http"

	"github.com/go-chi/chi/v5/middleware"

	"cloud-api/backend/common/logs"
	"cloud-api/backend/common/server"

	"github.com/go-chi/chi/v5"

	"cloud-api/backend/vnet-svc/config"
	"cloud-api/backend/vnet-svc/ports"
	"cloud-api/backend/vnet-svc/service"
)

func main() {
	cfg := config.GetConfig()
	logs.Init(cfg.IsLocalEnv)
	ctx := context.Background()
	// scpdestination.Destination.Get("terraform-agent")
	application := service.NewApplication(ctx)

	server.RunHTTPServer(cfg.ServerPort, func(router chi.Router) http.Handler {
		router.Use(middleware.Recoverer)
		return ports.HandlerWithOptions(
			ports.NewHttpServer(application),
			ports.ChiServerOptions{
				BaseRouter: router,
				// Middlewares: []ports.MiddlewareFunc{
				// 	auth.JwtAuthValidator(),
				// },
			},
		)
	})

}
