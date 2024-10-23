package service

import (
	"context"
	"log"
	"time"

	"github.com/Simo672K/issue-tracker/internal/db"
	"github.com/Simo672K/issue-tracker/internal/db/job"
	"github.com/Simo672K/issue-tracker/internal/db/model"
	"github.com/Simo672K/issue-tracker/internal/db/repository"
	"github.com/Simo672K/issue-tracker/utils"
	_ "github.com/lib/pq"
)

func CreateUser(user *model.User) error {
	errChan := make(chan error, 1)
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*200)
	defer cancel()

	db := db.GetDBConn()

	// hashes the password to create before creating a new user
	hpasswd, err := utils.HashPassword(user.HashedPassword)
	if err != nil {
		return err
	}

	// updates the user ported password to the hashed version
	user.HashedPassword = string(hpasswd)
	ur := repository.NewPGUserRepository(db)

	// creating the user and returning the user id
	userID, err := ur.Create(ctx, user)
	if err != nil {
		log.Fatal("Failed to create new user:", err)
		return err
	}

	// creating the profile using profile job go routine
	go job.CreateProfileJob(ctx, errChan, db, userID)

	if err := <-errChan; err != nil {
		log.Fatal("Failed to create user profile:", err)
		return err
	}

	return nil
}
