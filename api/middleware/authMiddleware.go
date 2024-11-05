package middleware

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/Simo672K/issue-tracker/utils"
	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware(ctx context.Context, handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-type", "application/json")
		cookie, err := r.Cookie("jwt_tokens")
		if err != nil {
			jsonErr := utils.HttpError().SetError(w, http.StatusForbidden, "FORBIDDEN", "Forbidden 403, cannot access.")
			w.Write(jsonErr)
			return
		}

		// extracting tokens from cookie
		accessToken, refreshToken := utils.GetTokensFromCookie(cookie)

		// validating tokens
		isAccessTokenValid, _ := utils.IsTokenValid(accessToken, utils.ACCESS_TOKEN)
		isRefreshTokenValid, _ := utils.IsTokenValid(refreshToken, utils.REFRESH_TOKEN)
		// extracting tokens payloads
		refreshTokenPayload, err := utils.ExtractTokenPayload(refreshToken, utils.REFRESH_TOKEN)
		if err != nil {
			utils.WriteJsonError(
				w,
				http.StatusBadRequest,
				"BAD_REQUEST",
				"An error has occured, please try later.",
			)
			return
		}

		accessTokenPayload, err := utils.ExtractTokenPayload(accessToken, utils.ACCESS_TOKEN)
		if err != nil {
			utils.WriteJsonError(
				w,
				http.StatusBadRequest,
				"INTERNAL_SERVER_ERROR",
				"An error has occured, please try later.",
			)
			return
		}
		ctx := context.WithValue(r.Context(), "profileId", (*accessTokenPayload)["sub"])
		r = r.WithContext(ctx)

		if isAccessTokenValid {
			handler.ServeHTTP(w, r)
			return
		}

		if isRefreshTokenValid {
			uid := (*refreshTokenPayload)["uid"]

			if (*refreshTokenPayload)["uid"] != (*accessTokenPayload)["uid"] {
				utils.WriteJsonError(
					w,
					http.StatusForbidden,
					"FORBIDDEN",
					"Access denied, invalid credentials",
				)
				return
			}

			// new access token
			accessPayload := jwt.MapClaims{}
			accessPayload["uid"] = uid
			accessPayload["email"] = (*accessTokenPayload)["email"]
			accessPayload["sub"] = (*accessTokenPayload)["sub"]
			accessPayload = utils.TokenPayloadConsruct(accessPayload, time.Minute*10)

			newTokens, err := utils.GenerateJwtTokens(accessPayload, uid.(string))
			if err != nil {

			}

			newTokens.RefreshToken = refreshToken

			// Setting up new cookie
			cookieVal := fmt.Sprintf("access_token:%s,refresh_token:%s", newTokens.AccessToken, newTokens.RefreshToken)

			// Overriding the cookie
			utils.SetTokenCookie(w, cookieVal)
			return
		}

		utils.WriteJsonError(
			w,
			http.StatusForbidden,
			"FORBIDDEN",
			"Access denied, invalid credentials",
		)
	})
}
