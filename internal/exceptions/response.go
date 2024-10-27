// internal/exceptions/response.go
package exceptions

import (
	"encoding/json"
	"net/http"
)

// Response represents the standard response format
type Response struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
}

// SendResponse sends a JSON response with the given status code
func SendResponse(w http.ResponseWriter, statusCode int, message string, data interface{}, err string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	response := Response{
		Message: message,
		Data:    data,
		Error:   err,
	}
	json.NewEncoder(w).Encode(response)
}
