package user

import "time"

// DAO represents a user
type DAO struct {
	ID        int       `json:"id"`
	Username  string    `json:"username"`
	Password  string    `json:"password_hash"`
	CreatedAt time.Time `json:"created_at"`
}
