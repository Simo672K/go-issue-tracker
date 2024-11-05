package utils

import "net/http"

func WriteJsonError(w http.ResponseWriter, status int, err, msg string) {
	jsonErr := HttpError().
		SetError(
			w,
			status,
			err,
			msg,
		)
	w.Header().Set("Content-type", "application/json")
	w.Write(jsonErr)
}
