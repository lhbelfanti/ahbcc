package corpus

import (
	"context"
	"fmt"

	"ahbcc/cmd/api/tweets"
	"ahbcc/cmd/api/tweets/categorized"
	"ahbcc/cmd/api/tweets/quotes"
	"ahbcc/internal/log"
)

// Create retrieves the information from the categorized_tweets table and inserts the tweets with all their information
// into the corpus table. It only considers the 'POSITIVE' and 'NEGATIVE' categorizations.
type Create func(ctx context.Context) error

// MakeCreate creates a new Create function
func MakeCreate(selectByCategorizations categorized.SelectByCategorizations, selectTweetByID tweets.SelectByID, selectTweetQuoteByID quotes.SelectByID, deleteAllCorpusRows DeleteAll, insertCorpusRow Insert) Create {
	var categorizations = []string{categorized.VerdictPositive, categorized.VerdictNegative}

	return func(ctx context.Context) error {
		categorizedTweets, err := selectByCategorizations(ctx, categorizations)
		if err != nil {
			log.Error(ctx, err.Error())
			return FailedToRetrieveCategorizedTweets
		}

		rows := make([]DTO, 0, len(categorizedTweets))
		for _, categorizedTweet := range categorizedTweets {
			tweetData, err := selectTweetByID(ctx, categorizedTweet.TweetID)
			if err != nil {
				log.Error(ctx, err.Error())
				continue
			}

			row := DTO{
				TweetAuthor:   tweetData.Author,
				TweetAvatar:   tweetData.Avatar,
				TweetText:     tweetData.TextContent,
				TweetImages:   tweetData.Images,
				IsTweetAReply: tweetData.IsAReply,
			}

			if tweetData.QuoteID != nil {
				tweetQuoteData, err := selectTweetQuoteByID(ctx, *tweetData.QuoteID)
				if err != nil {
					log.Error(ctx, err.Error())
					continue
				}

				row.QuoteAuthor = tweetQuoteData.Author
				row.QuoteAvatar = tweetQuoteData.Avatar
				row.QuoteText = tweetQuoteData.TextContent
				row.QuoteImages = tweetQuoteData.Images
				row.IsQuoteAReply = tweetQuoteData.IsAReply
			}

			rows = append(rows, row)
		}

		err = deleteAllCorpusRows(ctx)
		if err != nil {
			log.Error(ctx, err.Error())
			return FailedToCleanUpCorpusTable
		}

		var inserted int
		for _, row := range rows {
			err = insertCorpusRow(ctx, row)
			if err != nil {
				log.Error(ctx, err.Error())
				continue
			}

			inserted++
		}

		log.Info(ctx, fmt.Sprintf("Inserted %d/%d rows into the corpus table\n", inserted, len(rows)))

		return nil
	}
}
