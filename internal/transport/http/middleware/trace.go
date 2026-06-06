package middleware

import (
	core_logger "authTest/internal/platform/logger"
	http_response "authTest/internal/transport/http/response"
	"log/slog"
	"net/http"
	"time"
)

func Trace() Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			logger := core_logger.FromContext(ctx)
			newRW := http_response.NewResponseWriter(w)

			logger.Debug("incoming http request")
			t := time.Now()

			next.ServeHTTP(newRW, r)

			logger.Debug(
				"Finnished handling http request",
				slog.String("HTTP method", r.Method),
				slog.Duration("handling duration", time.Since(t)),
				slog.Int("status_code", newRW.GetStatusCode()),
			)
		})
	}
}
