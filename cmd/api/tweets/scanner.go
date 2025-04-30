package tweets

import (
	"github.com/jackc/pgx/v5"

	"ahbcc/cmd/api/tweets/quotes"
)

/*// MakeCollectTweetWithQuote is a custom CollectRows implementation to handle the information obtained from the
// SELECT used in the function SelectBySearchCriteriaIDYearAndMonth
func MakeCollectTweetWithQuote() database.CollectRows[TweetDTO] {
	return func(rows pgx.Rows) ([]TweetDTO, error) {
		return pgx.CollectRows(rows, func(row pgx.CollectableRow) (TweetDTO, error) {
			return ScanTweetWithQuote(row)
		})
	}
}*/

// CustomScanner is a custom scanner to parse the row obtained and return a TweetDTO which also contains a quotes.QuoteDTO
func CustomScanner() pgx.RowToFunc[TweetDTO] {
	return func(row pgx.CollectableRow) (TweetDTO, error) {
		var tweetDTO TweetDTO
		var quoteDTO quotes.QuoteDTO

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
			return TweetDTO{}, err
		}

		if tweetDTO.QuoteID != nil {
			tweetDTO.Quote = &quoteDTO
		}

		return tweetDTO, nil
	}
}
