package repository

import (
	"context"
	"database/sql"

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

func (ppr *PostgresProjectRepo) Create(ctx context.Context, project *model.Project) (string, error) {
	var projectId string
	sqlQuery := `INSERT INTO project (project_name) VALUES ($1) RETURNING id;`

	if err := ppr.DB.QueryRowContext(ctx, sqlQuery, project.ProjectName).Scan(&projectId); err != nil {
		return "", err
	}

	return projectId, nil
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
