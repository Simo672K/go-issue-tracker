package model

import "time"

type Project struct {
	Id              string
	ProjectName     string
	ProjectProgress float32
	CreatedAt       *time.Time
}

type ProjectOwner struct {
	Id        string
	OwnerId   string
	ProjectId string
}

type ProjectManager struct {
	Id        string
	ProfileID string
	ProjectId string
}

type ProjectDev struct {
	Id        string
	ProfileID string
	ProjectId string
}
