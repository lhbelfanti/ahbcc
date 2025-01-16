package response_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	"ahbcc/internal/http/response"
)

func TestSend_success(t *testing.T) {
	tests := []struct {
		name    string
		code    int
		message string
		data    interface{}
		err     interface{}
		want    response.DTO
	}{
		{
			name:    "Success response with data",
			code:    http.StatusOK,
			message: "Request successful",
			data:    map[string]string{"key": "value"},
			err:     nil,
			want: response.DTO{
				Code:    http.StatusOK,
				Message: "Request successful",
				Data:    map[string]string{"key": "value"},
				Error:   nil,
			},
		},
		{
			name:    "Error response with details",
			code:    http.StatusUnauthorized,
			message: "Unauthorized access",
			data:    nil,
			err:     map[string]string{"details": "Invalid token"},
			want: response.DTO{
				Code:    http.StatusUnauthorized,
				Message: "Unauthorized access",
				Data:    nil,
				Error:   map[string]string{"details": "Invalid token"},
			},
		},
		{
			name:    "Empty data response",
			code:    http.StatusNoContent,
			message: "No content",
			data:    nil,
			err:     nil,
			want: response.DTO{
				Code:    http.StatusNoContent,
				Message: "No content",
				Data:    nil,
				Error:   nil,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()

			response.Send(w, tt.code, tt.message, tt.data, tt.err)
			resp := w.Result()
			defer resp.Body.Close()

			var got response.DTO
			err := json.NewDecoder(resp.Body).Decode(&got)
			assert.NoError(t, err, "Failed to decode response")

			jsonWant, _ := json.Marshal(tt.want)
			jsonGot, _ := json.Marshal(got)
			assert.JSONEq(t, string(jsonWant), string(jsonGot), "Response mismatch")
		})
	}
}
