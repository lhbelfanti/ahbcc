package tweets

import (
	"github.com/jackc/pgx/v5"

	"ahbcc/cmd/api/tweets/quotes"
)

// CustomScanner is a custom scanner to parse the row retrieved and return a TweetDTO which also contains a quotes.QuoteDTO
func CustomScanner() pgx.RowToFunc[CustomTweetDTO] {
	return func(row pgx.CollectableRow) (CustomTweetDTO, error) {
		var tweetDTO CustomTweetDTO
		var quoteDTO quotes.CustomQuoteDTO

		err := row.Scan(
			&tweetDTO.ID,
			&tweetDTO.Author,
			&tweetDTO.Avatar,
			&tweetDTO.PostedAt,
			&tweetDTO.IsAReply,
			&tweetDTO.TextContent,
			&tweetDTO.Images,
			&tweetDTO.QuoteID,
			&tweetDTO.SearchCriteriaID,
			&quoteDTO.Author,
			&quoteDTO.Avatar,
			&quoteDTO.PostedAt,
			&quoteDTO.IsAReply,
			&quoteDTO.TextContent,
			&quoteDTO.Images,
		)
		if err != nil {
			return CustomTweetDTO{}, err
		}

		if tweetDTO.QuoteID != nil {
			tweetDTO.Quote = &quoteDTO
		}

		return tweetDTO, nil
	}
}
