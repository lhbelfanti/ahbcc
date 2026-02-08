package user

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"

	"github.com/lhbelfanti/corpus-creator/internal/database"
	"github.com/lhbelfanti/corpus-creator/internal/log"
)

type (
	// Exists validates if the given username was already registered in the database
	Exists func(ctx context.Context, username string) (bool, error)

	// SelectByUsername retrieves a user by its username
	SelectByUsername func(ctx context.Context, username string) (DAO, error)
)

// MakeExists creates a new Exists
func MakeExists(db database.Connection) Exists {
	const query string = `
		SELECT EXISTS (
			SELECT 1 
			FROM users 
			WHERE username = $1
		);
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

// MakeSelectByUsername creates a new SelectByUsername
func MakeSelectByUsername(db database.Connection) SelectByUsername {
	const query string = `
		SELECT id, username, password_hash, created_at
		FROM users
		WHERE username = $1;
	`

	return func(ctx context.Context, username string) (DAO, error) {
		var user DAO
		err := db.QueryRow(ctx, query, username).Scan(
			&user.ID,
			&user.Username,
			&user.PasswordHash,
			&user.CreatedAt,
		)
		if errors.Is(err, pgx.ErrNoRows) {
			log.Error(ctx, err.Error())
			return DAO{}, NoUserFoundForTheGivenUsername
		} else if err != nil {
			log.Error(ctx, err.Error())
			return DAO{}, FailedExecuteQueryToRetrieveUser
		}

		return user, nil
	}
}
