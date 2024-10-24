package utils

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JwtToken struct {
	AccessToken  string
	RefreshToken string
}

func GenerateJwtTokens(payload jwt.MapClaims) (*JwtToken, error) {
	accessSecret, refreshSecret := os.Getenv("ACCESS_TOKEN_SECRET"), os.Getenv("REFRESH_TOKEN_SECRET")
	token := JwtToken{
		AccessToken:  "",
		RefreshToken: "",
	}

	accToken := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	refToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"iat": time.Now().Unix(),
		"exp": time.Now().Add(time.Hour * 24 * 15).Unix(),
	})
	// signing the access token
	accessToken, err := accToken.SignedString([]byte(accessSecret))
	if err != nil {
		return nil, err
	}

	// signing the refresh token
	refreshToken, err := refToken.SignedString([]byte(refreshSecret))
	if err != nil {
		return nil, err
	}

	token.AccessToken = accessToken
	token.RefreshToken = refreshToken
	return &token, nil
}
