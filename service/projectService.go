package service

import (
	"context"
	"fmt"
	"log"

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
