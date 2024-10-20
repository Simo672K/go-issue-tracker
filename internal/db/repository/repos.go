package repository

import (
	"context"

	"github.com/Simo672K/issue-tracker/internal/db/model"
)

type UserRepository interface {
	Find(ctx context.Context, email string) (*model.User, error)
	FindAll(ctx context.Context) ([]*model.User, error)
	Create(ctx context.Context, user *model.User) error
	Update(ctx context.Context, user *model.User) error
	Delete(ctx context.Context, id string) error
}
