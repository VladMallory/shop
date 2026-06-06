package app

import (
	product_repository "authTest/internal/features/product/repository"
	product_service "authTest/internal/features/product/service"
	product_transport "authTest/internal/features/product/transport/http"
	database_postgres "authTest/internal/platform/db"
	core_logger "authTest/internal/platform/logger"
	http_server "authTest/internal/server"
	"authTest/internal/transport/http/middleware"
	"context"
	"os/signal"
	"syscall"
)

type App struct {
	httpServer *http_server.HTTPServer
}

func New() (*App, error) {
	ctx, cancel := signal.NotifyContext(
		context.Background(),
		syscall.SIGTERM,
		syscall.SIGINT,
	)
	defer cancel()

	// === LOGGER ===
	loggerCfg, err := core_logger.NewConfig()
	if err != nil {
		return nil, err
	}

	logger, err := core_logger.NewLogger(loggerCfg)
	if err != nil {
		return nil, err
	}
	defer logger.Close()

	pool, err := database_postgres.NewConnectionPool(ctx, database_postgres.NewConfigMust())
	if err != nil {
		return nil, err
	}

	productRepository := product_repository.NewProductRepository(pool)
	productService := product_service.NewProductService(productRepository)
	productHTTPHandler := product_transport.NewProductHTTPHandler(productService)

	httpServer := http_server.NewHTTPServer(
		http_server.NewConfigMust(),
		logger,
		middleware.RequestID(),
		middleware.Logger(logger),
		middleware.Trace(),
		middleware.Panic(),
	)

	apiVersionRouterV1 := http_server.NewAPIVersionRouter(http_server.APIVersionV1)
	apiVersionRouterV1.RegisterRoutes(productHTTPHandler.Routes()...)

	httpServer.RegisterRouters(apiVersionRouterV1)

	return &App{
		httpServer: httpServer,
	}, nil
}

func (a *App) Run() {
	ctx := context.Background()
	if err := a.httpServer.Run(ctx); err != nil {
		return
	}
}
