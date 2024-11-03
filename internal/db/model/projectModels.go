package model

import "time"

type Project struct {
	Id              string     `json:"projectId"`
	ProjectName     string     `json:"projectName"`
	ProjectProgress float32    `json:"progress"`
	CreatedAt       *time.Time `json:"createdAt"`
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
