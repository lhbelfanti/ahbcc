package user

// DTO represents a user to be inserted into the 'user' table
type DTO struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
