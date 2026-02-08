package auth

import (
	"context"
	"time"

	"golang.org/x/crypto/bcrypt"

	"github.com/lhbelfanti/corpus-creator/cmd/api/user"
	"github.com/lhbelfanti/corpus-creator/cmd/api/user/session"
	"github.com/lhbelfanti/corpus-creator/internal/log"
)

// LogIn logs the user in. It first verifies if the user exists in the database and then compares its password hash with
// a hashed version of the password given throw parameter, if they match, the user is allowed to log in.
// Lastly, it creates the user session and returns it with its expiration time.
type LogIn func(ctx context.Context, user user.DTO) (string, time.Time, error)

// MakeLogIn creates a new LogIn
func MakeLogIn(selectUserByUsername user.SelectByUsername, deleteExpiredUserSessions session.DeleteExpiredSessions, createSessionToken session.CreateToken) LogIn {
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

		err = deleteExpiredUserSessions(ctx, userDAO.ID)
		if err != nil {
			log.Warn(ctx, err.Error())
			// We don't want abort login due this cleanup
		}

		token, expiresAt, err := createSessionToken(ctx, userDAO.ID)
		if err != nil {
			log.Error(ctx, err.Error())
			return "", time.Time{}, FailedToCreateUserSession
		}

		return token, expiresAt, nil
	}
}
