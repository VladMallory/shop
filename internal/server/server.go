package http_server

import (
	core_logger "authTest/internal/logger"
	core_middleware "authTest/internal/middleware"
	"context"
	"errors"
	"fmt"
	"net/http"
)

type HTTPServer struct {
	mux    *http.ServeMux
	config Config
	logger *core_logger.Logger

	middleware []core_middleware.Middleware
}

func NewHTTPServer(
	config Config,
	logger *core_logger.Logger,
	middleware ...core_middleware.Middleware,
) *HTTPServer {
	return &HTTPServer{
		mux:        http.NewServeMux(),
		config:     config,
		logger:     logger,
		middleware: middleware,
	}
}

func (s *HTTPServer) RegisterRouters(routers ...*APIVersionRouter) {
	for _, router := range routers {
		prefix := "/api/" + string(router.apiVersion)

		r := router.WithMiddleware()
		s.mux.Handle(prefix+"/", http.StripPrefix(prefix, r))
	}
}

func (s *HTTPServer) Run(ctx context.Context) error {
	mux := core_middleware.ChainMiddleware(s.mux, s.middleware...)
	server := &http.Server{
		Addr:    s.config.Address,
		Handler: mux,
	}

	ch := make(chan error, 1)
	go func() {
		defer close(ch)
		s.logger.Warn("starting http server")

		err := server.ListenAndServe()
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			ch <- err
		}
	}()

	select {
	case err := <-ch:
		if err != nil {
			return fmt.Errorf("run server error: %w", err)
		}
	case <-ctx.Done():
		s.logger.Warn("shutting down http server")
		
		shutdownCtx, cancel := context.WithTimeout(context.Background(), s.config.ShutdownTimeout)
		defer cancel()

		if err := server.Shutdown(shutdownCtx); err != nil {
			_ = server.Close()
			return fmt.Errorf("shutdown server: %w", err)
		}
	}

	return nil
}
