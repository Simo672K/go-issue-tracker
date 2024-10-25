package repository

import (
	"context"

	"github.com/Simo672K/issue-tracker/internal/db/model"
)

type UserRepository interface {
	Find(ctx context.Context, email string) (*model.User, error)
	FindAll(ctx context.Context) ([]*model.User, error)
	Create(ctx context.Context, user *model.User) (string, error)
	Update(ctx context.Context, user *model.User) error
	Delete(ctx context.Context, id string) error
}

type ProfileRepository interface {
	Create(ctx context.Context, profile *model.Profile) error
}

type ProjectRepository interface {
	Find(ctx context.Context, id string) (*model.Project, error)
	FindAll(ctx context.Context, ownerId string) ([]*model.Project, error)
	Create(ctx context.Context, project *model.Project, ownerId string) error
	Update(ctx context.Context, project *model.Project) error
	Delete(ctx context.Context, id string) error
}
