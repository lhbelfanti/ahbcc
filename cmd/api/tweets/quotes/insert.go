package quotes

import (
	"context"
	"errors"
	"time"

	"github.com/jackc/pgx/v5"

	"github.com/lhbelfanti/corpus-creator/internal/database"
	"github.com/lhbelfanti/corpus-creator/internal/log"
)

// InsertSingle inserts a new QuoteDTO into 'quotes' table and returns the PK
type InsertSingle func(ctx context.Context, quote *QuoteDTO) (int, error)

// MakeInsertSingle creates a new InsertSingle
func MakeInsertSingle(db database.Connection) InsertSingle {
	const query string = `
			INSERT INTO tweets_quotes(is_a_reply, author, avatar, posted_at, text_content, images) 
			VALUES ($1, $2, $3, $4, $5, $6)
			RETURNING id;
		`

	return func(ctx context.Context, quote *QuoteDTO) (int, error) {
		if quote == nil {
			return -1, NothingToInsertWhenQuoteIsNil
		}

		var postedAt *time.Time
		if quote.PostedAt != "" {
			parsedDate, err := time.Parse(time.RFC3339, quote.PostedAt)
			if err != nil {
				log.Warn(ctx, err.Error())
			} else {
				postedAt = &parsedDate
			}
		}

		var quoteID int
		err := db.QueryRow(ctx, query, quote.IsAReply, quote.Author, quote.Avatar, postedAt, quote.TextContent, quote.Images).Scan(&quoteID)
		if errors.Is(err, pgx.ErrNoRows) {
			log.Error(ctx, err.Error())
			return -1, FailedToInsertQuote
		}

		return quoteID, nil
	}
}
