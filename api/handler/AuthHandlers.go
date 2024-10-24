package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Simo672K/issue-tracker/internal/db"
	"github.com/Simo672K/issue-tracker/internal/db/model"
	"github.com/Simo672K/issue-tracker/internal/db/repository"
	"github.com/Simo672K/issue-tracker/service"
)

// signup and registration handler
func AuthSignUpHandler(w http.ResponseWriter, r *http.Request) {
	var user model.User
	db := db.GetDBConn()
	ur := repository.NewPGUserRepository(db)

	// Binding request body data with user model
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		log.Fatal("An error accured while parsing json data", err)
		return
	}

	// handling user creation with user service
	if err := service.CreateUser(&user, ur, db); err != nil {
		log.Fatal("An error accured while creating user:", err)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

// TODO : signin handler - should accept credentials
// TODO : -> pass it to an auth-service -> verify credentials -> return a jwt or invalid credentials code error
func AuthSignInHandler(w http.ResponseWriter, r *http.Request) {
	var credentials service.Credentials
	db := db.GetDBConn()
	ur := repository.NewPGUserRepository(db)
	header := w.Header()

	if err := json.NewDecoder(r.Body).Decode(&credentials); err != nil {
		log.Fatal("An error accured while parsing credentials data", err)
		return
	}

	header.Add("Content-Type", "application/json")
	tokens, status := service.SignInService(credentials, ur)

	switch status {
	case http.StatusOK:
		tokens, err := json.Marshal(*tokens)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Write(tokens)
	default:
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
}

// TODO : change password - update the current password to a new one
func AuthChangePassword() {}

// TODO : reset password - when password forgoted
func AuthResetPassword() {}
