package utils

import (
	"encoding/json"
	"net/http"
)

type JsonHttpErrorResponse struct {
	Status  int    `json:"status"`
	Error   string `json:"error"`
	Message string `json:"message"`
}

func HttpError() *JsonHttpErrorResponse {
	return &JsonHttpErrorResponse{}
}

func (jher *JsonHttpErrorResponse) SetError(w http.ResponseWriter, status int, err, msg string) []byte {
	jher.Status = status
	jher.Error = err
	jher.Message = msg
	w.WriteHeader(status)
	jsonErr, _ := json.Marshal(jher)
	return jsonErr
}

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
