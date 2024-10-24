package service

import (
	"context"

	"github.com/Simo672K/issue-tracker/internal/db/repository"
	"github.com/Simo672K/issue-tracker/utils"
)

type Credentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func SignInService(credentials Credentials, ur repository.UserRepository) error {
	ctx := context.Background()
	user, err := ur.Find(ctx, credentials.Email)

	if err != nil {
		return err
	}

	// checks if user's credentials are legit
	if utils.IsCredentialValid(user.HashedPassword, credentials.Password) {

	}

	return nil
}
