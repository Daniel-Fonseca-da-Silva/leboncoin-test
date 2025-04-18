package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Response struct {
	Message    string      `json:"message,omitempty"`
	ErrorCode  int         `json:"errorCode,omitempty"`
	Data       interface{} `json:"data,omitempty"`
	StatusCode int         `json:"-"`
}

func SendError(w http.ResponseWriter, code int, msg string) {
	response := Response{
		Message:    msg,
		ErrorCode:  code,
		StatusCode: code,
	}
	sendJSON(w, response)
}

func SendSuccess(w http.ResponseWriter, operation string, data interface{}) {
	response := Response{
		Message:    fmt.Sprintf("Operation %s successful", operation),
		Data:       data,
		StatusCode: http.StatusOK,
	}
	sendJSON(w, response)
}

func sendJSON(w http.ResponseWriter, response Response) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(response.StatusCode)
	json.NewEncoder(w).Encode(response)
}
