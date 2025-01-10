package users

import (
	"context"

	"ahbcc/internal/database"
)

// Insert inserts a new UserDTO into 'users' table
type Insert func(ctx context.Context, user UserDTO) error

// MakeInsert creates a new Insert
func MakeInsert(db database.Connection) Insert {
	const query string = `
		INSERT INTO users(username, password_hash) 
		VALUES ($1, $2)
		ON CONFLICT (username, password_hash) DO NOTHING;
	`

	return func(ctx context.Context, user UserDTO) error {
		_, err := db.Exec(ctx, query, user.Username, user.Password)
		if err != nil {
			return FailedToInsertUser
		}

		return nil
	}
}
