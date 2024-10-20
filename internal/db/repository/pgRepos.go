package repository

import (
	"context"
	"database/sql"
	"time"

	"github.com/Simo672K/issue-tracker/internal/db/model"
	_ "github.com/lib/pq"
)

type PostgresUserRepo struct {
	DB *sql.DB
}

// find user based on it's email
func (pur *PostgresUserRepo) Find(ctx context.Context, email string) (*model.User, error) {
	var user model.User

	sqlQuery := "SELECT * FROM user WHERE email=$1"
	ctx, cancel := context.WithTimeout(ctx, time.Millisecond*200)
	defer cancel()

	if err := pur.DB.QueryRowContext(
		ctx,
		sqlQuery,
		email).Scan(
		&user.Email,
		&user.HashedPassword,
		&user.Name,
		&user.UserID); err != nil {
		return nil, err
	}

	return &user, nil
}

func (pur *PostgresUserRepo) FindAll(ctx context.Context) ([]*model.User, error) {
	return nil, nil
}

func (pur *PostgresUserRepo) Create(ctx context.Context, user *model.User) error {
	return nil
}

func (pur *PostgresUserRepo) Update(ctx context.Context, user *model.User) error {
	return nil

}

func (pur *PostgresUserRepo) Delete(ctx context.Context, id string) error {
	return nil
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &PostgresUserRepo{
		DB: db,
	}
}
