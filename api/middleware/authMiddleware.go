package middleware

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/Simo672K/issue-tracker/utils"
)

func AuthMiddleware(ctx context.Context, handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("jwt_tokens")
		if err != nil {
			w.Write([]byte("wrong"))
			http.Error(w, "something went wrong", http.StatusUnauthorized)
			return
		}

		// extracting tokens from cookie
		tokens := strings.Split(cookie.Value, ",")
		accessToken := strings.Replace(tokens[0], "access_token:", "", 1)
		// refreshToken := strings.Replace(tokens[1], "refresh_token:", "", 1)

		isValid := utils.IsCredentialValid(accessToken, "ACCESS_TOKEN")
		fmt.Println(isValid)
		handler.ServeHTTP(w, r)
	})
}
