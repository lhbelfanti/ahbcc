package tweets

import (
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"

	"github.com/lhbelfanti/corpus-creator/cmd/api/tweets/quotes"
)

// CustomScanner is a custom scanner to parse the row retrieved and return a TweetDTO which also contains a quotes.QuoteDTO
func CustomScanner() pgx.RowToFunc[CustomTweetDTO] {
	return func(row pgx.CollectableRow) (CustomTweetDTO, error) {
		var (
			tweetDTO CustomTweetDTO

			// Nullable variables for scanning
			quoteAuthor      pgtype.Text
			quoteAvatar      pgtype.Text
			quotePostedAt    pgtype.Timestamp
			quoteIsAReply    pgtype.Bool
			quoteTextContent pgtype.Text
			quoteImages      []string
		)

		err := row.Scan(
			&tweetDTO.ID,
			&tweetDTO.StatusID,
			&tweetDTO.Author,
			&tweetDTO.Avatar,
			&tweetDTO.PostedAt,
			&tweetDTO.IsAReply,
			&tweetDTO.TextContent,
			&tweetDTO.Images,
			&tweetDTO.QuoteID,
			&tweetDTO.SearchCriteriaID,
			&quoteAuthor,
			&quoteAvatar,
			&quotePostedAt,
			&quoteIsAReply,
			&quoteTextContent,
			&quoteImages,
		)
		if err != nil {
			return CustomTweetDTO{}, err
		}

		if tweetDTO.QuoteID != nil {
			tweetDTO.Quote = &quotes.CustomQuoteDTO{
				Author:      quoteAuthor.String,
				Avatar:      pgTextToStringPtr(quoteAvatar),
				PostedAt:    quotePostedAt.Time,
				IsAReply:    quoteIsAReply.Bool,
				TextContent: pgTextToStringPtr(quoteTextContent),
				Images:      quoteImages,
			}
		}

		return tweetDTO, nil
	}
}

func pgTextToStringPtr(text pgtype.Text) *string {
	if text.Valid && text.String != "" {
		return &text.String
	}

	return nil
}
