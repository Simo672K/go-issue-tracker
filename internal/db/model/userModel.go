package model

import (
	"time"
)

type User struct {
	UserID         string
	Name           string
	Email          string
	HashedPassword string
	Created        *time.Time
}
