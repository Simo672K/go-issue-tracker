package handler

import (
	"net/http"

	"github.com/Simo672K/issue-tracker/service"
)

func VerifyUserEmail(w http.ResponseWriter, r *http.Request) {
	verifId := r.PathValue("verificationId")

	if err := service.ValidateVerification(verifId); err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Write([]byte("User verified!"))
}

func NewVerificationHandler(w http.ResponseWriter, r *http.Request) {
	service.SendVerificationEmail()
}
