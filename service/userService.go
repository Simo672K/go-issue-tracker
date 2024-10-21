package service

import (
	"context"
	"database/sql"
	"time"

	"github.com/Simo672K/issue-tracker/internal/db"
	"github.com/Simo672K/issue-tracker/internal/db/model"
	"github.com/Simo672K/issue-tracker/internal/db/repository"
	"github.com/Simo672K/issue-tracker/utils"
	_ "github.com/lib/pq"
)

func CreateUser(user *model.User) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*200)
	defer cancel()

	connStr := db.GetDBConnStr()
	db, err := sql.Open("postgres", connStr)

	if err != nil {
		return err
	}

	// hashes the password to create before creating a new user
	hpasswd, err := utils.HashPassword(user.HashedPassword)
	if err != nil {
		return err
	}

	// updates the user ported password to the hashed version
	user.HashedPassword = string(hpasswd)

	ur := repository.NewPGUserRepository(db)
	ur.Create(ctx, user)
	return nil
}

func createProfile() {}
