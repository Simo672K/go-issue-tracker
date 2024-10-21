package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Simo672K/issue-tracker/internal/db/model"
	"github.com/Simo672K/issue-tracker/service"
)

func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	var user model.User

	// Binding request body data with user model
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		log.Fatal("An error accured while parsing json data", err)
		return
	}

	// handling user creation with user service
	if err := service.CreateUser(&user); err != nil {
		log.Fatal("An error accured while creating user:", err)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
