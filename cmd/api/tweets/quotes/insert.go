package quotes

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
	"log/slog"

	"ahbcc/internal/database"
)

// InsertSingle inserts a new QuoteDTO into 'quotes' table and returns the PK
type InsertSingle func(quote QuoteDTO) (int, error)

// MakeInsertSingle creates a new InsertSingle
func MakeInsertSingle(db database.Connection) InsertSingle {
	const query string = `
			INSERT INTO tweets_quotes(is_a_reply, has_text, has_images, text_content, images) 
			VALUES ($1, $2, $3, $4, $5)
			RETURNING id
		`

	return func(quote QuoteDTO) (int, error) {
		var quoteID int

		err := db.QueryRow(context.Background(), query, quote.IsAReply, quote.TextContent, quote.Images).Scan(&quoteID)
		if errors.Is(err, pgx.ErrNoRows) {
			slog.Error(err.Error())
			return -1, FailedToInsertQuote
		}

		return quoteID, nil
	}
}
