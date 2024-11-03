package repository

import (
	"context"
	"database/sql"
	"log"
)

type PostgresProjectOwnerRepo struct {
	DB *sql.DB
}

func NewPGProjectOwnerRepository(db *sql.DB) ProjectOwnerRepository {
	return &PostgresProjectOwnerRepo{
		DB: db,
	}
}

func (ppor *PostgresProjectOwnerRepo) Create(ctx context.Context, profileId, projectId string) error {
	sqlQuery := `INSERT INTO project_owner (owner_id, project_id) VALUES ($1, $2);`

	if _, err := ppor.DB.ExecContext(ctx, sqlQuery, profileId, projectId); err != nil {
		log.Fatalf("Failed to create project of owner #%s ", profileId)
		return err
	}

	return nil
}
