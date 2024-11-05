package repository

import (
	"context"
	"database/sql"
	"log"

	"github.com/Simo672K/issue-tracker/internal/db/model"
)

type PostgresProfileRepo struct {
	DB *sql.DB
}

func (ppr *PostgresProfileRepo) Create(ctx context.Context, profile *model.Profile) error {
	sqlQuery := "INSERT INTO profile (user_id, username) VALUES ($1, $2)"
	if _, err := ppr.DB.ExecContext(ctx, sqlQuery, profile.UserID, profile.Username); err != nil {
		log.Fatalf("Failed to create userId #%s profile's", profile.UserID)
		return err
	}
	return nil
}

func (ppr *PostgresProfileRepo) FindByUserId(ctx context.Context, userId string) (*model.Profile, error) {
	sqlQuery := `SELECT * FROM profile WHERE user_id = $1`
	var profile model.Profile

	if err := ppr.DB.QueryRowContext(ctx, sqlQuery, userId).Scan(&profile.Id, &profile.UserID, &profile.Username, &profile.Created); err != nil {
		log.Fatalf("Failed to get user #%s profile", userId)
		return nil, err
	}
	return &profile, nil
}

func NewPGProfileRepo(db *sql.DB) ProfileRepository {
	return &PostgresProfileRepo{
		DB: db,
	}
}
