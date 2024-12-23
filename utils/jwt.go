package utils

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JwtToken struct {
	AccessToken  string
	RefreshToken string
}

type Token struct {
	payload jwt.MapClaims
}

const (
	ACCESS_TOKEN  = "ACCESS_TOKEN"
	REFRESH_TOKEN = "REFRESH_TOKEN"
)

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

func tokenParser(token, tokenType string) (*jwt.Token, error) {
	var secret string

	switch tokenType {
	case ACCESS_TOKEN:
		secret = accessSecret
	case REFRESH_TOKEN:
		secret = refreshSecret
	}

	jwtToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	// if err != nil {
	// 	return nil
	// }
	return jwtToken, err
}

func IsTokenValid(token string, tokenType string) (bool, error) {
	jwtToken, err := tokenParser(token, tokenType)
	// fmt.Println(jwtToken)
	if err != nil {
		return jwtToken.Valid, err
	}
	return jwtToken.Valid, nil
}

func ExtractTokenPayload(token string, tokenType string) (*jwt.MapClaims, error) {
	jwtToken, err := tokenParser(token, tokenType)
	payload := jwtToken.Claims.(jwt.MapClaims)

	if err != nil {
		return &payload, fmt.Errorf("failed to extract token payload: %s", err)
	}
	return &payload, nil
}

func TokenPayloadConsruct(payload jwt.MapClaims, duration time.Duration) jwt.MapClaims {
	payload["iat"] = time.Now().Unix()
	payload["exp"] = time.Now().Add(duration).Unix()
	return payload
}

func AccessTokenPayloadConstructor(id string, email, profileId string) jwt.MapClaims {
	payload := jwt.MapClaims{
		"uid":   id,
		"email": email,
		"sub":   profileId,
	}
	payload = TokenPayloadConsruct(payload, time.Minute*10)
	return payload
}

func GetTokensFromCookie(jwtCookie *http.Cookie) (string, string) {
	tokens := strings.Split(jwtCookie.Value, ",")
	accessToken := strings.Replace(tokens[0], "access_token:", "", 1)
	refreshToken := strings.Replace(tokens[1], "refresh_token:", "", 1)

	return accessToken, refreshToken
}

func NewToken() *Token {
	return &Token{
		payload: jwt.MapClaims{},
	}
}
func (t *Token) Add(key string, value interface{}) {
	t.payload[key] = value
}

func (t *Token) Construct(duration time.Duration) jwt.MapClaims {
	t.payload = TokenPayloadConsruct(t.payload, duration)

	return t.payload
}
