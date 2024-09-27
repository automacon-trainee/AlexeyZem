package middleware

import (
	"context"
	"net/http"
	"strings"

	"project/internal/API/gRPCAuth"
)

type Middleware struct {
	auth gRPCAuth.AuthServiceClient
}

func NewMiddleware(auth gRPCAuth.AuthServiceClient) *Middleware {
	return &Middleware{auth: auth}
}

func (m *Middleware) AuthVerify(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenStr := r.Header.Get("Authorization")
		tokenStr = strings.TrimPrefix(tokenStr, "Bearer ")
		token := gRPCAuth.Token{Token: tokenStr}
		_, err := m.auth.VerifyToken(context.Background(), &token)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			_, _ = w.Write([]byte(err.Error()))
		} else {
			next.ServeHTTP(w, r)
		}
	})
}
