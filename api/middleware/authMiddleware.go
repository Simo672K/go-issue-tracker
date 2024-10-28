package middleware

import (
	"context"
	"fmt"
	"net/http"
)

func AuthMiddleware(ctx context.Context, handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("jwt_tokens")
		if err != nil {
			w.Write([]byte("wrong"))
			http.Error(w, "something went wrong", http.StatusUnauthorized)
			return
		}
		fmt.Println(cookie)
		handler.ServeHTTP(w, r)
	})
}
