package users

// MockUserDTO mocks UserDTO
func MockUserDTO() UserDTO {
	return UserDTO{
		Username: "test@test.com",
		Password: "password",
	}
}
