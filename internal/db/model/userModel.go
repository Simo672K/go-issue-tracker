package model

import (
	"time"
)

type User struct {
	UserID         string     `json:"userId"`
	Name           string     `json:"name"`
	Email          string     `json:"email"`
	HashedPassword string     `json:"password"`
	Created        *time.Time `json:"created"`
}
