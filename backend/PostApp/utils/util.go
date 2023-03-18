package utils

import (
	"encoding/json"
	"net/http"
)

func Message(success bool, message string, data interface{}) map[string]interface{} {
	return map[string]interface{}{
		"success": success,
		"message": message,
		"data":    data,
	}
}

func Respond(w http.ResponseWriter, data interface{}, status int) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}
