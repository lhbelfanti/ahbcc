package auth

import (
	"context"
	"time"

	"github.com/lhbelfanti/corpus-creator/cmd/api/user"
)

// MockSignUp mocks SignUp function
func MockSignUp(err error) SignUp {
	return func(ctx context.Context, user user.DTO) error {
		return err
	}
}

// MockLogIn mocks LogIn function
func MockLogIn(token string, expiresAt time.Time, err error) LogIn {
	return func(ctx context.Context, user user.DTO) (string, time.Time, error) {
		return token, expiresAt, err
	}
}

// MockLogout mocks LogOut function
func MockLogout(err error) LogOut {
	return func(ctx context.Context, user string) error {
		return err
	}
}
