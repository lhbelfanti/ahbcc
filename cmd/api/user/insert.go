package user

import (
	"context"

	"ahbcc/internal/database"
	"ahbcc/internal/log"
)

// Insert inserts a new DTO into 'user' table
type Insert func(ctx context.Context, user DTO) error

// MakeInsert creates a new Insert
func MakeInsert(db database.Connection) Insert {
	const query string = `
		INSERT INTO users(username, password_hash) 
		VALUES ($1, $2)
		ON CONFLICT (username, password_hash) DO NOTHING;
	`

	return func(ctx context.Context, user DTO) error {
		_, err := db.Exec(ctx, query, user.Username, user.Password)
		if err != nil {
			log.Error(ctx, err.Error())
			return FailedToInsertUser
		}

		return nil
	}
}
