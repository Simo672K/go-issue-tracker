package utils

import (
	"net/http"
	"time"
)

func SetTokenCookie(w http.ResponseWriter, value string) {
	cookie := &http.Cookie{
		Name:     "jwt_tokens",
		Value:    value,
		Expires:  time.Now().Add(time.Hour * 24 * 15),
		HttpOnly: true,
		Path:     "/",
		SameSite: http.SameSiteStrictMode,
	}

	http.SetCookie(w, cookie)
}
