package middleware

import (
	core_logger "authTest/internal/platform/logger"
	"log/slog"
	"net/http"
)

func Logger(logger *core_logger.Logger) Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			requestID := r.Header.Get(requestIDHeader)

			logger = logger.With(
				slog.String("requestID", requestID),
				slog.String("url", r.URL.String()),
			)

			ctx := core_logger.ContextWithLogger(r.Context(), logger)

			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
