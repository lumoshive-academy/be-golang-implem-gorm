package middleware

import (
	"context"
	"net/http"

	"go.uber.org/zap"
)

type AuthMiddleware struct {
	Logger *zap.Logger
}

func NewAuthMiddleware(logger *zap.Logger) AuthMiddleware {
	return AuthMiddleware{logger}
}

func (m *AuthMiddleware) Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), "userid", 123)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
