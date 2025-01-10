package users

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"

	"ahbcc/internal/database"
	"ahbcc/internal/log"
)

// UserExists validates if the given username was already registered in the database
type UserExists func(ctx context.Context, username string) (bool, error)

// MakeUserExists creates a new UserExists
func MakeUserExists(db database.Connection) UserExists {
	const query string = `
		SELECT EXISTS (
			SELECT 1 
			FROM users 
			WHERE username = $1
		)
	`

	return func(ctx context.Context, username string) (bool, error) {
		var applied bool

		err := db.QueryRow(ctx, query, username).Scan(&applied)
		if errors.Is(err, pgx.ErrNoRows) {
			log.Error(ctx, err.Error())
			return false, FailedToRetrieveIfUserAlreadyExists
		}

		return applied, nil
	}
}
