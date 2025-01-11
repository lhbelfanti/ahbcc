package users

import "context"

// MockUserExists mocks UserExists function
func MockUserExists(userExists bool, err error) UserExists {
	return func(ctx context.Context, username string) (bool, error) {
		return userExists, err
	}
}

// MockInsert mocks Insert function
func MockInsert(err error) Insert {
	return func(ctx context.Context, user UserDTO) error {
		return err
	}
}

// MockUserDTO mocks UserDTO
func MockUserDTO() UserDTO {
	return UserDTO{
		Username: "test@test.com",
		Password: "password",
	}
}
