package service

import (
	"context"
	"net/http"

	"github.com/Simo672K/issue-tracker/internal/db/repository"
	"github.com/Simo672K/issue-tracker/utils"
)

type Credentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func SignInService(credentials Credentials, ur repository.UserRepository, pr repository.ProfileRepository) (*utils.JwtToken, int) {

	ctx := context.Background()
	user, err := ur.Find(ctx, credentials.Email)

	if err != nil {
		return nil, http.StatusUnauthorized
	}

	profile, err := pr.FindByUserId(ctx, user.Id)
	if err != nil {
		return nil, http.StatusUnauthorized
	}

	// checks if user's credentials are legit
	if utils.IsCredentialValid(user.HashedPassword, credentials.Password) {
		id := utils.StrUniqueId()

		payload := utils.AccessTokenPayloadConstructor(id, user.Email, profile.Id)

		// Generating jwt tokens
		token, err := utils.GenerateJwtTokens(payload, id)

		if err != nil {
			return nil, http.StatusUnauthorized
		}

		return token, http.StatusOK
	}

	return nil, http.StatusUnauthorized
}
