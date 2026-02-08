package session

import (
	"context"

	"github.com/lhbelfanti/corpus-creator/internal/database"
	"github.com/lhbelfanti/corpus-creator/internal/log"
)

// Insert inserts a new session DAO into 'user_sessions' table
type Insert func(ctx context.Context, session DAO) error

// MakeInsert creates a new Insert function
func MakeInsert(db database.Connection) Insert {
	const query string = `
		INSERT INTO users_sessions(user_id, token, expires_at) 
		VALUES ($1, $2, $3)
	`

	return func(ctx context.Context, session DAO) error {
		_, err := db.Exec(ctx, query, session.UserID, session.Token, session.ExpiresAt)
		if err != nil {
			log.Error(ctx, err.Error())
			return FailedToInsertUserSession
		}

		return nil
	}
}
