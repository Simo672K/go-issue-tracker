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
	var project model.Project

	sqlQuery := `SELECT * FROM project WHERE id=$1`

	if err := ppr.DB.QueryRowContext(ctx, sqlQuery, id).
		Scan(&project.Id, &project.ProjectName, &project.ProjectProgress, &project.CreatedAt); err != nil {
		return nil, err
	}

	return &project, nil
}

func (ppr *PostgresProjectRepo) FindAll(ctx context.Context, profileId string) ([]model.Project, error) {
	var projects []model.Project

	sqlQuery := `
	SELECT 
		p.id, 
		p.project_name, 
		p.project_progress, 
		p.created_at
	FROM project p
	LEFT JOIN project_owner po ON p.id = po.project_id AND po.owner_id = $1
	LEFT JOIN project_manager pm ON p.id = pm.project_id AND pm.profile_id = $1
	LEFT JOIN project_dev pd ON p.id = pd.project_id AND pd.profile_id = $1
	WHERE po.owner_id IS NOT NULL OR pm.profile_id IS NOT NULL OR pd.profile_id IS NOT NULL;
	`

	rows, err := ppr.DB.QueryContext(ctx, sqlQuery, profileId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var project model.Project

		if err := rows.Scan(&project.Id, &project.ProjectName, &project.ProjectProgress, &project.CreatedAt); err != nil {
			return nil, err
		}

		projects = append(projects, project)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return projects, nil
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
