package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/Simo672K/issue-tracker/internal/db"
	"github.com/Simo672K/issue-tracker/internal/db/model"
	"github.com/Simo672K/issue-tracker/internal/db/repository"
	"github.com/Simo672K/issue-tracker/service"
	"github.com/Simo672K/issue-tracker/utils"
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
	resMessage := make(map[string]string)
	var credentials service.Credentials
	db := db.GetDBConn()
	ur := repository.NewPGUserRepository(db)

	// Decoding request body and binding it to credentials
	if err := json.NewDecoder(r.Body).Decode(&credentials); err != nil {
		log.Fatal("An error accured while parsing credentials data", err)
		return
	}

	// setting up tokens
	tokens, status := service.SignInService(credentials, ur)

	switch status {
	case http.StatusOK:

		cookieVal := fmt.Sprintf("access_token:%s,refresh_token:%s", tokens.AccessToken, tokens.RefreshToken)

		//  Setting tokens as an httponly cookie
		utils.SetTokenCookie(w, string(cookieVal))
		resMessage["message"] = "Signed in successfully!"
		jsonMsg, _ := json.Marshal(resMessage)
		w.Write(jsonMsg)
	default:
		w.WriteHeader(http.StatusUnauthorized)
		resMessage["error"] = "Email or password incorrect"
		jsonMsg, _ := json.Marshal(resMessage)
		w.Write(jsonMsg)
	}
}

// TODO : change password - update the current password to a new one
func AuthChangePassword(w http.ResponseWriter, r *http.Request) {}

// TODO : reset password - when password forgoted
func AuthResetPassword(w http.ResponseWriter, r *http.Request) {}

func AuthTest(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("jwt_tokens")
	if err != nil {
		w.Write([]byte("wrong"))
		http.Error(w, "something went wrong", http.StatusUnauthorized)
		return
	}

	fmt.Println(cookie.Value)
}
