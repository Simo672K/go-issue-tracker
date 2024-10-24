package service

import (
	"context"
	"net/http"
	"time"

	"github.com/Simo672K/issue-tracker/internal/db/repository"
	"github.com/Simo672K/issue-tracker/utils"
	"github.com/golang-jwt/jwt/v5"
)

type Credentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func SignInService(credentials Credentials, ur repository.UserRepository) (*utils.JwtToken, int) {

	ctx := context.Background()
	user, err := ur.Find(ctx, credentials.Email)

	if err != nil {
		return nil, http.StatusUnauthorized
	}

	// checks if user's credentials are legit
	if utils.IsCredentialValid(user.HashedPassword, credentials.Password) {
		payload := jwt.MapClaims{
			"email": user.Email,
			"sub":   user.Id,
			"iat":   time.Now().Unix(),
			"exp":   time.Now().Add(time.Minute * 10).Unix(),
		}

		// Generating jwt tokens
		token, err := utils.GenerateJwtTokens(payload)

		if err != nil {
			return nil, http.StatusUnauthorized
		}

		return token, http.StatusOK
	}

	return nil, http.StatusUnauthorized
}
