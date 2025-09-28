package utils

import (
	"encoding/json"
	"flutter-chat/dto"
	"net/http"
)

// Respond with success (data can be any struct or nil)
func RespondSuccess(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	resp := dto.APIResponse{
		Code: code,
		Data: data,
	}
	json.NewEncoder(w).Encode(resp)
}

// Respond with error (errMsg is a string describing the error)
func RespondError(w http.ResponseWriter, errData *dto.ErrorData) {
	code := errData.Code
	if code == 0 {
		code = http.StatusInternalServerError // Default to 500 if not set
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	resp := dto.APIResponse{
		Code:  code,
		Error: errData,
	}
	json.NewEncoder(w).Encode(resp)
}
