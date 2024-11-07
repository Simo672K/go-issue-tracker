package auth

import (
	"context"
	"fmt"

	"github.com/Simo672K/issue-tracker/internal/db"
	"github.com/Simo672K/issue-tracker/internal/db/repository"
)

type Role int

const (
	OWNER Role = iota
	MANAGER
	DEVELOPER
)

type Permission struct {
	role Role
}

func NewPermission(role Role) *Permission {
	return &Permission{
		role,
	}
}

func (p *Permission) HasAccessTo(ctx context.Context, profileId, projectId string) bool {
	pgdb := db.GetDBConn()
	switch p.role {
	case OWNER:
		por := repository.NewPGProjectOwnerRepository(pgdb)
		isProjectOwner, err := por.IsProjectOwner(ctx, profileId, projectId)
		if err != nil {
			fmt.Println(err)
			return false
		}
		return isProjectOwner

	case MANAGER:
		return false

	case DEVELOPER:
		return false

	default:
		return false
	}
}
