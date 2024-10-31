package middleware

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/Simo672K/issue-tracker/utils"
	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware(ctx context.Context, handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-type", "application/json")
		errMessage := make(map[string]string)
		cookie, err := r.Cookie("jwt_tokens")
		if err != nil {
			errMessage["message"] = "403 - Access denied!"
			jsonMsg, _ := utils.JsonStringfiedHttpResponse(w, errMessage)
			http.Error(w, jsonMsg, http.StatusForbidden)
			return
		}

		// extracting tokens from cookie
		tokens := strings.Split(cookie.Value, ",")
		accessToken := strings.Replace(tokens[0], "access_token:", "", 1)
		refreshToken := strings.Replace(tokens[1], "refresh_token:", "", 1)

		// validating tokens
		isAccessTokenValid, _ := utils.IsTokenValid(accessToken, utils.ACCESS_TOKEN)
		isRefreshTokenValid, _ := utils.IsTokenValid(refreshToken, utils.REFRESH_TOKEN)

		if isAccessTokenValid {
			handler.ServeHTTP(w, r)
			return
		}

		if isRefreshTokenValid {
			// extracting tokens payloads
			refreshTokenPayload, err := utils.ExtractTokenPayload(refreshToken, utils.REFRESH_TOKEN)
			accessTokenPayload, _ := utils.ExtractTokenPayload(accessToken, utils.ACCESS_TOKEN)
			uid := (*refreshTokenPayload)["uid"]

			if err != nil {
				errMessage["message"] = "An error has accured"
				jsonMsg, _ := utils.JsonStringfiedHttpResponse(w, errMessage)
				http.Error(w, jsonMsg, http.StatusUnauthorized)
				return
			}

			if (*refreshTokenPayload)["uid"] != (*accessTokenPayload)["uid"] {
				errMessage["message"] = "Invalid credentials"
				jsonMsg, _ := utils.JsonStringfiedHttpResponse(w, errMessage)
				http.Error(w, jsonMsg, http.StatusUnauthorized)
				return
			}

			// new access token
			accessPayload := jwt.MapClaims{}
			accessPayload["uid"] = uid
			accessPayload["email"] = (*accessTokenPayload)["email"]
			accessPayload["sub"] = (*accessTokenPayload)["sub"]
			accessPayload = utils.TokenPayloadConsruct(accessPayload, time.Minute*10)

			newTokens, err := utils.GenerateJwtTokens(accessPayload, uid.(string))
			newTokens.RefreshToken = refreshToken

			// Setting up new cookie
			cookieVal := fmt.Sprintf("access_token:%s,refresh_token:%s", newTokens.AccessToken, newTokens.RefreshToken)

			// Overriding the cookie
			utils.SetTokenCookie(w, cookieVal)
			return
		}

		errMessage["message"] = "Invalid credentials"
		jsonMsg, _ := utils.JsonStringfiedHttpResponse(w, errMessage)
		http.Error(w, jsonMsg, http.StatusUnauthorized)
	})
}
