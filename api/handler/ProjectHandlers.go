package handler

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/Simo672K/issue-tracker/internal/db"
	"github.com/Simo672K/issue-tracker/internal/db/model"
	"github.com/Simo672K/issue-tracker/internal/db/repository"
	"github.com/Simo672K/issue-tracker/service"
	"github.com/Simo672K/issue-tracker/utils"
)

func CreateProjectHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-type", "application/json")
	resMsg := utils.NewJsonMsg()
	profileId := r.URL.Query().Get("profile_id")
	var project model.Project

	if err := json.NewDecoder(r.Body).Decode(&project); err != nil {
		log.Fatal("An error accured while parsing json data", err)
		jsonErr := utils.HttpError().
			SetError(
				w,
				http.StatusInternalServerError,
				err.Error(),
				"Failed to parse data, try again later.",
			)

		w.Write(jsonErr)
		return
	}

	db := db.GetDBConn()
	ctx := context.Background()

	projectRepo := repository.NewPGProjectRepository(db)
	projectOwnerRepo := repository.NewPGProjectOwnerRepository(db)

	if err := service.CreateNewProjectService(
		ctx,
		projectRepo,
		projectOwnerRepo,
		&project,
		profileId,
	); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	resMsg.Add("message", "project created successfully!")
	successResp, _ := resMsg.ToHttpResponse()
	w.Write(successResp)
}
