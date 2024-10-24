package job

import (
	"context"
	"database/sql"

	"github.com/Simo672K/issue-tracker/internal/db/model"
	"github.com/Simo672K/issue-tracker/internal/db/repository"
)

func CreateProfileJob(ctx context.Context, errChan chan error, db *sql.DB, userId string) {
	profile := &model.Profile{
		UserID: userId,
	}

	ppr := repository.NewPGProfileRepo(db)
	if err := ppr.Create(ctx, profile); err != nil {
		errChan <- err
		return
	}

	errChan <- nil
}
