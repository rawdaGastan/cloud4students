// Package routes for API endpoints
package routes

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/rawdaGastan/cloud4students/internal"
	"github.com/rawdaGastan/cloud4students/models"
)

// Router struct holds db model and configurations
type Router struct {
	config *internal.Configuration
	db     models.DB
}

// NewRouter create new router with db
func NewRouter(config internal.Configuration, db models.DB) (r Router) {
	return Router{&config, db}
}

// ErrorMsg holds errors
type ErrorMsg struct {
	Error string `json:"err"`
}

// ResponseMsg holds messages and needed data
type ResponseMsg struct {
	Message string      `json:"msg"`
	Data    interface{} `json:"data,omitempty"`
}

// writeErrResponse write error messages in api
func writeErrResponse(w http.ResponseWriter, err error) {
	jsonErrRes, _ := json.Marshal(ErrorMsg{Error: err.Error()})
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)
	_, err = w.Write(jsonErrRes)
	if err != nil {
		log.Printf("write error response failed %v", err.Error())
	}
}

// writeNotFoundResponse write error messages in api
func writeNotFoundResponse(w http.ResponseWriter, err error) {
	jsonErrRes, _ := json.Marshal(ErrorMsg{Error: err.Error()})
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound)
	_, err = w.Write(jsonErrRes)
	if err != nil {
		log.Printf("write not found error response failed %v", err.Error())
	}
}

// writeMsgResponse write response messages for api
func writeMsgResponse(w http.ResponseWriter, message string, data interface{}) {
	contentJSON, err := json.Marshal(ResponseMsg{Message: message, Data: data})
	if err != nil {
		writeErrResponse(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(contentJSON)
	if err != nil {
		writeErrResponse(w, fmt.Errorf("write message response failed %v", err))
	}
}
