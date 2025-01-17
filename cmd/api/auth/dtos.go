package auth

import "time"

// LoginResponseDTO represents the response of the LogIn endpoint
type LoginResponseDTO struct {
	Token     string    `json:"token"`
	ExpiresAt time.Time `json:"expires_at"`
}
