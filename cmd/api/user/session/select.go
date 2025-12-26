package session

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
	
	"ahbcc/internal/database"
	"ahbcc/internal/log"
)

type SelectUserIDByToken func(ctx context.Context, token string) (int, error)

func MakeSelectUserIDByToken(db database.Connection) SelectUserIDByToken {
	const query string = `
		SELECT user_id
		FROM users_sessions
		WHERE token = $1;
	`

	return func(ctx context.Context, token string) (int, error) {
		var userID int
		err := db.QueryRow(ctx, query, token).Scan(&userID)
		if errors.Is(err, pgx.ErrNoRows) {
			log.Error(ctx, err.Error())
			return 0, NoUserIDFoundForTheGivenToken
		} else if err != nil {
			log.Error(ctx, err.Error())
			return 0, FailedToExecuteQueryToRetrieveUserID
		}

		return userID, nil
	}
}
