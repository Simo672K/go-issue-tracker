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

/*
-- get if profile is manager
SELECT EXISTS (

	SELECT 1
	FROM project_manager
	WHERE profile_id = 'YOUR_PROFILE_ID'
	  AND project_id = 'YOUR_PROJECT_ID'

);

-- get if profile is developer
SELECT EXISTS (

	SELECT 1
	FROM project_dev
	WHERE profile_id = 'YOUR_PROFILE_ID'
	  AND project_id = 'YOUR_PROJECT_ID'

);
*/

func (ppor *PostgresProjectOwnerRepo) IsProjectOwner(ctx context.Context, profileId, projectId string) (bool, error) {
	sqlQuery := `
	SELECT EXISTS (
		SELECT 1
		FROM project_owner
		WHERE owner_id = $1
			AND project_id = $2
	);
	`
	var answer bool
	if err := ppor.DB.QueryRowContext(ctx, sqlQuery, profileId, projectId).Scan(&answer); err != nil {
		return false, err
	}
	return answer, nil
}
