package response

import (
	"encoding/json"
	"net/http"
)

// Send writes a standardized JSON response to the client.
// It accepts an HTTP status code, a message, optional data, and optional error details.
// The response format includes the code, message, and either data or error information.
func Send(w http.ResponseWriter, code int, message string, data interface{}, err interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	resp := DTO{
		Code:    code,
		Message: message,
		Data:    data,
		Error:   err,
	}

	_ = json.NewEncoder(w).Encode(resp)
}
