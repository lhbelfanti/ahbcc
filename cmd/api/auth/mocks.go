package auth

import (
	"context"
	
	"ahbcc/cmd/api/users"
)

// MockSignUp mocks SignUp function
func MockSignUp(err error) SignUp {
	return func(ctx context.Context, user users.UserDTO) error {
		return err
	}
}
