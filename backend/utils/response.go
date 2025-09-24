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
func RespondError(w http.ResponseWriter, code int, message string, fields ...dto.FieldError) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	errData := &dto.ErrorData{
		Message: message,
	}
	if len(fields) > 0 {
		errData.Fields = fields
	}
	resp := dto.APIResponse{
		Code:  code,
		Error: errData,
	}
	json.NewEncoder(w).Encode(resp)
}
