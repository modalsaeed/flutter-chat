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
		Success: true,
		Code:    code,
		Data:    data,
	}
	json.NewEncoder(w).Encode(resp)
}

// Respond with error (errMsg is a string describing the error)
func RespondError(w http.ResponseWriter, code int, errMsg string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	resp := dto.APIResponse{
		Success: false,
		Code:    code,
		Error:   errMsg,
	}
	json.NewEncoder(w).Encode(resp)
}
