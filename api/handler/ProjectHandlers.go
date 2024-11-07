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
	successResp, _ := resMsg.ToJson()
	w.Write(successResp)
}

func ListAssociatedProjectsHandler(w http.ResponseWriter, r *http.Request) {
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
	projectsList := utils.NewJsonMsg()
	projectsList.Add("projects", projects)
	projectsList.Add("count", len(projects))

	jsonResp, err := projectsList.ToJson()
	if err != nil {
		log.Println(err)
		utils.WriteJsonError(
			w,
			http.StatusBadRequest,
			err.Error(),
			"An error has occured, try again later.",
		)
	}

	w.Write(jsonResp)
}

func GetProjectInfoHandler(w http.ResponseWriter, r *http.Request) {
	projectId := r.PathValue("projectId")
	w.Header().Add("Content-type", "application/json")
	resMsg := utils.NewJsonMsg()

	db := db.GetDBConn()
	projectRepo := repository.NewPGProjectRepository(db)
	project, err := service.GetProjectInfoService(r.Context(), projectRepo, projectId)
	if err != nil {
		log.Println(err)
		utils.WriteJsonError(
			w,
			http.StatusBadRequest,
			err.Error(),
			"An error has occured internaly, try again later",
		)
		return
	}
	resMsg.Add("data", project)
	jsonMsg, err := resMsg.ToJson()
	if err != nil {
		log.Println(err)
		utils.WriteJsonError(
			w,
			http.StatusBadRequest,
			err.Error(),
			"An error has occured internaly, try again later",
		)
	}

	w.Write(jsonMsg)
}

func DeleteProjectHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("OK"))
}
