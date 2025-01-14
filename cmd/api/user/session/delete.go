package session

import (
	"context"

	"ahbcc/internal/database"
	"ahbcc/internal/log"
)

type (
	// Delete deletes a session, seeking it by its token
	Delete func(ctx context.Context, token string) error

	// DeleteUserExpiredSessions deletes the expired sessions of a given user
	DeleteUserExpiredSessions func(ctx context.Context, userID int) error
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

// MakeDeleteUserExpiredSessions creates a new DeleteUserExpiredSessions
func MakeDeleteUserExpiredSessions(db database.Connection) DeleteUserExpiredSessions {
	const query string = `
		DELETE FROM users_sessions
		WHERE user_id = $1
	`

	return func(ctx context.Context, userID int) error {
		_, err := db.Exec(ctx, query, userID)
		if err != nil {
			log.Error(ctx, err.Error())
			return FailedToDeleteExpiredUserSessions
		}

		return nil
	}
}
