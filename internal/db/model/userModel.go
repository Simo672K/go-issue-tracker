package model

import (
	"time"
)

type User struct {
	Id             string     `json:"userId"`
	Name           string     `json:"name"`
	Email          string     `json:"email"`
	HashedPassword string     `json:"password"`
	Created        *time.Time `json:"created"`
}

type Profile struct {
	Id       string     `json:"profileId"`
	UserID   string     `json:"userId"`
	Username string     `json:"username"`
	Created  *time.Time `json:"created"`
}
