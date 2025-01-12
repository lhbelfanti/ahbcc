package auth

import (
	"context"

	"ahbcc/cmd/api/user"
)

// MockSignUp mocks SignUp function
func MockSignUp(err error) SignUp {
	return func(ctx context.Context, user user.DTO) error {
		return err
	}
}
