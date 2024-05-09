package error_handler

import (
	"encoding/json"
	"net/http"
)

type ErrorResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Error   string `json:"error"`
}

func InvalidIdError(w http.ResponseWriter, message string) {
	errResponse := new(ErrorResponse)
	errResponse.Status = http.StatusBadRequest
	errResponse.Error = "Invalid id"
	errResponse.Message = message

	response, _ := json.Marshal(&errResponse)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest)
	w.Write(response)
}

func NotFoundError(w http.ResponseWriter, message string) {
	errResponse := new(ErrorResponse)
	errResponse.Status = http.StatusNotFound
	errResponse.Error = "not found"
	errResponse.Message = message

	response, _ := json.Marshal(&errResponse)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound)
	w.Write(response)
}

func BadRequestError(w http.ResponseWriter, message string) {
	errResponse := new(ErrorResponse)
	errResponse.Status = http.StatusBadRequest
	errResponse.Error = "bad request"
	errResponse.Message = message

	response, _ := json.Marshal(&errResponse)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest)
	w.Write(response)
}
