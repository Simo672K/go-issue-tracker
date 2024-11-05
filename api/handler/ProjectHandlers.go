package handler

import (
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
		log.Println("An error accured while parsing json data", err)
		jsonErr := utils.HttpError().
			SetError(
				w,
				http.StatusBadRequest,
				err.Error(),
				"Failed to parse data, try again later.",
			)

		w.Write(jsonErr)
		return
	}

	db := db.GetDBConn()
	ctx := r.Context()

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

func ListAssociatedProjectsHandler(w http.ResponseWriter, r *http.Request) {
	response := make(map[string]interface{})
	pgdb := db.GetDBConn()
	pr := repository.NewPGProjectRepository(pgdb)

	profileId := r.Context().Value("profileId")
	projects, err := service.ListAssociatedProjectsService(r.Context(), pr, profileId.(string))
	if err != nil {
		utils.WriteJsonError(
			w,
			http.StatusBadRequest,
			err.Error(),
			"An error has occured internaly, try again later",
		)
		return
	}
	response["projects"] = projects
	response["count"] = len(projects)

	resp, err := json.Marshal(response)
	w.Write(resp)
}
