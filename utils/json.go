package utils

import (
	"encoding/json"
	"net/http"
)

func JsonStringfiedHttpResponse(w http.ResponseWriter, msg interface{}) (string, error) {
	jsonMsg, err := json.Marshal(msg)

	if err != nil {
		http.Error(w, "An error has accured!", http.StatusInternalServerError)
		return "", err
	}

	return string(jsonMsg), nil
}
