package session

import (
	"context"

	"ahbcc/internal/database"
	"ahbcc/internal/log"
)

// Delete deletes a session
type Delete func(ctx context.Context, token string) error

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
