package middleware

import (
	"authTest/internal/service"
	"context"
	"net/http"
	"strings"
)

// userKey тип для ключа контекста, чтобы избежать коллизий.
type userKey struct{}

func UserIDFromCtx(ctx context.Context) int64 {
	return ctx.Value(userKey{}).(int64)
}

// Auth проверяет JWT в заголовке Authorization и кладёт user_id в контекст.
func Auth(authService *service.AuthService) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Достаём токен из заголовка Authorization: Bearer <token>
			raw := r.Header.Get("Authorization")
			tokenStr := strings.TrimPrefix(raw, "Bearer ")
			if tokenStr == "" || tokenStr == raw {
				http.Error(w, `{"error":"missing token"}`, http.StatusUnauthorized)

				return
			}

			// Валидируем токен через AuthService
			userID, err := authService.ValidateToken(tokenStr)
			if err != nil {
				http.Error(w, `{"error":"invalid token"}`, http.StatusUnauthorized)

				return
			}

			// Кладём userID в контекст и передаём дальше
			ctx := context.WithValue(r.Context(), userKey{}, userID)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
