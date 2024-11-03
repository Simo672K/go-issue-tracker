package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/Simo672K/issue-tracker/internal/db"
	"github.com/Simo672K/issue-tracker/internal/db/model"
	"github.com/Simo672K/issue-tracker/internal/db/repository"
	"github.com/Simo672K/issue-tracker/utils"
)

func CreateProject(w http.ResponseWriter, r *http.Request) {
	errMsg := utils.NewJsonMsg()
	successMsg := utils.NewJsonMsg()
	profileId := r.URL.Query().Get("profile_id")
	var project model.Project

	if err := json.NewDecoder(r.Body).Decode(&project); err != nil {
		errMsg.Add("message", "failed to parse data")
		errMsg.Add("error", err)
		strErrMsg, _ := errMsg.ToString()

		log.Fatal("An error accured while parsing json data", err)
		http.Error(w, strErrMsg, http.StatusInternalServerError)
		return
	}

	db := db.GetDBConn()
	ctx := context.Background()

	projectRepo := repository.NewPGProjectRepository(db)
	projectOwnerRepo := repository.NewPGProjectOwnerRepository(db)

	projectId, err := projectRepo.Create(ctx, &project)
	if err != nil {
		errMsg.Add("message", "failed to create new project")
		errMsg.Add("error", err)
		strErrMsg, _ := errMsg.ToString()

		log.Fatal("An error accured while creating project", err)
		http.Error(w, strErrMsg, http.StatusInternalServerError)
		return
	}

	fmt.Println(projectId)
	if err := projectOwnerRepo.Create(ctx, profileId, projectId); err != nil {
		errMsg.Add("message", "operation failed")
		errMsg.Add("error", err)
		strErrMsg, _ := errMsg.ToString()

		log.Fatal("An error accured while creating project", err)
		http.Error(w, strErrMsg, http.StatusInternalServerError)
		return
	}

	successMsg.Add("message", "project created successfully!")
	successResp, _ := successMsg.ToHttpResponse()
	w.Write(successResp)
}
