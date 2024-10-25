package repository

import (
	"context"
	"database/sql"
	"log"

	"github.com/Simo672K/issue-tracker/internal/db/model"
)

type PostgresProjectRepo struct {
	DB *sql.DB
}

// find user based on it's email
func (ppr *PostgresProjectRepo) Find(ctx context.Context, id string) (*model.Project, error) {
	return nil, nil
}

func (ppr *PostgresProjectRepo) FindAll(ctx context.Context, ownerId string) ([]*model.Project, error) {
	return nil, nil
}

func (ppr *PostgresProjectRepo) Create(ctx context.Context, project *model.Project, ownerId string) error {
	sqlQuery := `
	BEGIN;
	INSERT INTO project (project_name) values ($1) RETURNING id AS new_project_id;
	INSERT INTO project_owner (owner_id, project_id) values (new_project_id, $2)
	COMMIT;
	`
	if _, err := ppr.DB.ExecContext(ctx, sqlQuery, project.ProjectName, ownerId); err != nil {
		log.Fatalf("Failed to create project of owner #%s ", ownerId)
		return err
	}

	return nil
}

func (ppr *PostgresProjectRepo) Update(ctx context.Context, project *model.Project) error {
	return nil

}

func (ppr *PostgresProjectRepo) Delete(ctx context.Context, id string) error {
	return nil
}

func NewPGProjectRepository(db *sql.DB) ProjectRepository {
	return &PostgresProjectRepo{
		DB: db,
	}
}
