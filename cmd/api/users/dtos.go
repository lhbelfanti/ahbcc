package users

// UserDTO represents a user to be inserted into the 'users' table
type UserDTO struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
