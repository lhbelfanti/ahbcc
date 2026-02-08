package session

import (
	"context"

	"github.com/lhbelfanti/corpus-creator/internal/database"
	"github.com/lhbelfanti/corpus-creator/internal/log"
)

type (
	// Delete deletes a session, seeking it by its token
	Delete func(ctx context.Context, token string) error

	// DeleteExpiredSessions deletes the expired sessions of a given user
	DeleteExpiredSessions func(ctx context.Context, userID int) error
)

// MakeDelete creates a new Delete
func MakeDelete(db database.Connection) Delete {
	const query string = `
		DELETE FROM users_sessions
		WHERE token = $1
	`

	return func(ctx context.Context, token string) error {
		_, err := db.Exec(ctx, query, token)
		if err != nil {
			log.Error(ctx, err.Error())
			return FailedToDeleteUserSession
		}

		return nil
	}
}

// MakeDeleteExpiredSessions creates a new DeleteExpiredSessions
func MakeDeleteExpiredSessions(db database.Connection) DeleteExpiredSessions {
	const query string = `
		DELETE FROM users_sessions
		WHERE user_id = $1 
		AND expires_at < NOW()
	`

	return func(ctx context.Context, userID int) error {
		_, err := db.Exec(ctx, query, userID)
		if err != nil {
			log.Error(ctx, err.Error())
			return FailedToDeleteExpiredSessions
		}

		return nil
	}
}
