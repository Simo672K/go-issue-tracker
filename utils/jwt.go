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

var (
	accessSecret  string
	refreshSecret string
)

func init() {
	accessSecret, refreshSecret = os.Getenv("ACCESS_TOKEN_SECRET"), os.Getenv("REFRESH_TOKEN_SECRET")
}

func GenerateJwtTokens(payload jwt.MapClaims, id string) (*JwtToken, error) {
	token := JwtToken{
		AccessToken:  "",
		RefreshToken: "",
	}

	// signing the access token
	accessToken, err := SignJwtToken(payload, []byte(accessSecret))
	if err != nil {
		return nil, err
	}

	// signing the refresh token
	refreshToken, err := SignJwtToken(jwt.MapClaims{
		"uid": id,
		"iat": time.Now().Unix(),
		"exp": time.Now().Add(time.Hour * 24 * 15).Unix(),
	}, []byte(refreshSecret))

	if err != nil {
		return nil, err
	}

	token.AccessToken = accessToken
	token.RefreshToken = refreshToken
	return &token, nil
}

func SignJwtToken(payload jwt.MapClaims, secret []byte) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	stringToken, err := token.SignedString([]byte(secret))

	if err != nil {
		return "", err
	}

	return stringToken, nil
}

func IsTokenValid(tokens JwtToken, tokenType string) bool {
	var secret string

	switch tokenType {
	case "ACCESS_TOKEN":
		secret = accessSecret
	case "REFRESH_TOKEN":
		secret = refreshSecret
	}

	jwtToken, err := jwt.Parse(tokens.AccessToken, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv(secret)), nil
	})
	if err != nil {
		return false
	}
	return jwtToken.Valid
}
