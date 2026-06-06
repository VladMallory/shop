package app

import (
	product_repository "authTest/internal/features/product/repository"
	product_service "authTest/internal/features/product/service"
	product_transport "authTest/internal/features/product/transport"
	database_postgres "authTest/internal/infrastructure/db"
	core_logger "authTest/internal/logger"
	"authTest/internal/middleware"
	http_server "authTest/internal/server"
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func UpdatedRun() {
	fmt.Println("1")
	loggerCfg := core_logger.NewConfigMust()
	ctx, cancel := signal.NotifyContext(
		context.Background(),
		syscall.SIGTERM,
		syscall.SIGINT,
	)

	defer cancel()

	logger, err := core_logger.NewLogger(loggerCfg)
	if err != nil {
		fmt.Println("https://www.youtube.com/shorts/WR0Uh3-AVNA")
		os.Exit(1)
	}

	defer logger.Close()
	logger.Debug("initializing connection pool")

	pool, err := database_postgres.NewConnectionPool(ctx, database_postgres.NewConfigMust())
	if err != nil {
		logger.Error("failed to initialize connection pool")
		os.Exit(1)
	}

	logger.Debug("initializing product feature")
	productRepository := product_repository.NewProductRepository(pool)
	productService := product_service.NewProductService(productRepository)
	productHTTPHandler := product_transport.NewProductHTTPHandler(productService)

	logger.Debug("initializing http server")
	httpServer := http_server.NewHTTPServer(
		http_server.NewConfigMust(),
		logger,
		middleware.RequestID(),
		middleware.Logger(logger),
		middleware.Trace(),
		middleware.Panic(),
	)

	logger.Debug("registering routes")
	apiVersionRouterV1 := http_server.NewAPIVersionRouter(http_server.APIVersionV1)
	apiVersionRouterV1.RegisterRoutes(productHTTPHandler.Routes()...)

	httpServer.RegisterRouters(apiVersionRouterV1)

	if err := httpServer.Run(ctx); err != nil {
		logger.Error("HTTP server run error: %w", err)
		os.Exit(1)
	}
}
