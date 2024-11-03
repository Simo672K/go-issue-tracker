package utils

import (
	"encoding/json"
	"net/http"
)

type JsonMessage struct {
	Body map[string]any
}

func NewJsonMsg() *JsonMessage {
	return &JsonMessage{
		Body: make(map[string]any),
	}
}

func (jm *JsonMessage) Add(key string, value any) {
	jm.Body[key] = value
}

func (jm *JsonMessage) ToString() (string, error) {
	rawMarshaledMsg, err := json.Marshal(jm.Body)
	if err != nil {
		return "", err
	}
	return string(rawMarshaledMsg), nil
}

func (jm *JsonMessage) ToHttpResponse() ([]byte, error) {
	jsonResp, err := json.Marshal(jm.Body)
	if err != nil {
		return nil, err
	}

	return jsonResp, nil
}

func JsonStringfiedHttpResponse(w http.ResponseWriter, msg interface{}) (string, error) {
	jsonMsg, err := json.Marshal(msg)

	if err != nil {
		http.Error(w, "An error has accured!", http.StatusInternalServerError)
		return "", err
	}

	return string(jsonMsg), nil
}
