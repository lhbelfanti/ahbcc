package auth

import (
	"context"

	"github.com/lhbelfanti/corpus-creator/cmd/api/user/session"
	"github.com/lhbelfanti/corpus-creator/internal/log"
)

// LogOut Logs the user out. It deletes the current user session associated to that user by the token
type LogOut func(ctx context.Context, token string) error

// MakeLogOut creates a new LogOut
func MakeLogOut(deleteUserSession session.Delete) LogOut {
	return func(ctx context.Context, token string) error {

		err := deleteUserSession(ctx, token)
		if err != nil {
			log.Error(ctx, err.Error())
			return FailedToDeleteUserSession
		}

		return nil
	}
}
