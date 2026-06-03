package middleware

import (
	core_logger "authTest/internal/logger"
	http_response "authTest/internal/transport/http/response"
	"net/http"
)

func Panic() Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			logger := core_logger.FromContext(r.Context())
			responseHandler := http_response.NewHTTPResponseHandler(logger, w)

			defer func() {
				if p := recover(); p != nil {
					responseHandler.RespondPanic(p, "got unexpected panic")
				}
			}()

			next.ServeHTTP(w, r)
		})
	}
}
