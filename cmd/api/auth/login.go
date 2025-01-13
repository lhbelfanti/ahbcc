package auth

import (
	"context"
	"golang.org/x/crypto/bcrypt"
	"time"

	"ahbcc/cmd/api/user"
	"ahbcc/cmd/api/user/session"
	"ahbcc/internal/log"
)

// LogIn logs the user in
type LogIn func(ctx context.Context, user user.DTO) (string, time.Time, error)

// MakeLogIn creates a new LogIn function
func MakeLogIn(selectUserByUsername user.SelectByUsername, createSessionToken session.CreateToken) LogIn {
	return func(ctx context.Context, user user.DTO) (string, time.Time, error) {
		userDAO, err := selectUserByUsername(ctx, user.Username)
		if err != nil {
			log.Error(ctx, err.Error())
			return "", time.Time{}, FailedToSelectUserByUsername
		}

		err = bcrypt.CompareHashAndPassword([]byte(userDAO.PasswordHash), []byte(user.Password))
		if err != nil {
			log.Error(ctx, err.Error())
			return "", time.Time{}, FailedToLoginDueWrongPassword
		}

		token, expiresAt, err := createSessionToken(ctx, userDAO.ID)
		if err != nil {
			log.Error(ctx, err.Error())
			return "", time.Time{}, FailedToCreateUserSession
		}

		return token, expiresAt, nil
	}
}
