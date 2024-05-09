package controllers

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Status int         `json:"status"`
	Data   interface{} `json:"data"`
	Meta   interface{} `json:"meta"`
}

func StatusOkResponse(w http.ResponseWriter, data interface{}, meta interface{}) {
	response := new(Response)
	response.Status = http.StatusOK
	response.Meta = meta
	response.Data = data

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	responseData, _ := json.Marshal(&response)
	w.Write([]byte(responseData))
}

func StatusCreatedResponse(w http.ResponseWriter, data interface{}, meta interface{}) {
	response := new(Response)
	response.Status = http.StatusCreated
	response.Meta = meta
	response.Data = data

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	responseData, _ := json.Marshal(&response)
	w.Write([]byte(responseData))
}
