package auth

import "time"

// LoginResponse represents the response of the LogIn endpoint
type LoginResponse struct {
	Token     string    `json:"token"`
	ExpiresAt time.Time `json:"expires_at"`
}
