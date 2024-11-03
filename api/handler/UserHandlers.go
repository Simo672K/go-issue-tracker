package handler

import (
	"fmt"
	"net/http"

	"github.com/Simo672K/issue-tracker/service"
)

func VerifyUserEmail(w http.ResponseWriter, r *http.Request) {
	userId := r.PathValue("userId")
	token := r.URL.Query().Get("token")

	fmt.Println(userId, token)
	// if err := service.ValidateVerification(verifId); err != nil {
	// 	http.Error(w, err.Error(), http.StatusNotFound)
	// 	return
	// }

	// w.Write([]byte("User verified!"))
}

func NewVerificationHandler(w http.ResponseWriter, r *http.Request) {
	service.SendVerificationEmail()
}
