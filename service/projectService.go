package service

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/Simo672K/issue-tracker/internal/db/model"
	"github.com/Simo672K/issue-tracker/internal/db/repository"
	"github.com/Simo672K/issue-tracker/utils"
)

func CreateNewProjectService(ctx context.Context, pr repository.ProjectRepository, por repository.ProjectOwnerRepository, project *model.Project, profileId string) error {
	errMsg := utils.NewJsonMsg()
	projectId, err := pr.Create(ctx, project)
	if err != nil {
		errMsg.Add("message", "failed to create new project")
		errMsg.Add("error", err)
		strErrMsg, _ := errMsg.ToString()

		log.Fatal("An error accured while creating project", err)
		return fmt.Errorf(strErrMsg)
	}

	if err := por.Create(ctx, profileId, projectId); err != nil {
		errMsg.Add("message", "operation failed")
		errMsg.Add("error", err)
		strErrMsg, _ := errMsg.ToString()

		log.Fatal("An error accured while creating project", err)
		return fmt.Errorf(strErrMsg)
	}
	return nil
}

func ListAssociatedProjectsService(ctx context.Context, pr repository.ProjectRepository, profileId string) ([]model.Project, error) {
	ctxTimout, cancel := context.WithTimeout(ctx, time.Millisecond*150)
	defer cancel()

	projects, err := pr.FindAll(ctxTimout, profileId)
	if err != nil {
		return nil, err
	}
	return projects, nil
}

func GetProjectInfoService(ctx context.Context, pr repository.ProjectRepository, projectId string) (*model.Project, error) {
	ctxTimeout, cancel := context.WithTimeout(ctx, time.Millisecond*150)
	defer cancel()

	project, err := pr.Find(ctxTimeout, projectId)
	if err != nil {
		return nil, err
	}
	return project, nil
}
