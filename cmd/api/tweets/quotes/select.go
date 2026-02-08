package quotes

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"

	"github.com/lhbelfanti/corpus-creator/internal/database"
	"github.com/lhbelfanti/corpus-creator/internal/log"
)

// SelectByID retrieves a tweet DAO by its ID
type SelectByID func(ctx context.Context, id int) (DAO, error)

// MakeSelectByID creates a new SelectByID function
func MakeSelectByID(db database.Connection) SelectByID {
	const query string = `SELECT tq.id, tq.author, tq.avatar, tq.posted_at, tq.is_a_reply, tq.text_content, tq.images
						  FROM tweets_quotes AS tq
						  WHERE tq.id = $1`

	return func(ctx context.Context, id int) (DAO, error) {
		var tweet DAO
		err := db.QueryRow(ctx, query, id).Scan(
			&tweet.ID,
			&tweet.Author,
			&tweet.Avatar,
			&tweet.PostedAt,
			&tweet.IsAReply,
			&tweet.TextContent,
			&tweet.Images,
		)
		if errors.Is(err, pgx.ErrNoRows) {
			log.Error(ctx, err.Error())
			return DAO{}, NoTweetQuoteFoundForTheGivenTweetQuoteID
		} else if err != nil {
			log.Error(ctx, err.Error())
			return DAO{}, FailedExecuteQueryToRetrieveTweetQuoteData
		}

		return tweet, nil
	}
}
