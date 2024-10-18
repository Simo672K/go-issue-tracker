package repository

import (
	"context"
	"database/sql"
	"log"

	"github.com/Simo672K/issue-tracker/internal/db/model"
)

type PostgresUserRepo struct {
	DB *sql.DB
}

func (pur *PostgresUserRepo) Find(ctx context.Context, id string) (*model.User, error) {
	return nil, nil
}

func (pur *PostgresUserRepo) FindAll(ctx context.Context, fieldName, value string) ([]*model.User, error) {
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

func Test() {
	// user := model.User{}
	connStr := "postgres://user:password@localhost/dbname?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	ctx := context.Background()
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	userRepo := NewUserRepository(db)
	userRepo.Find(ctx, "id")
}
