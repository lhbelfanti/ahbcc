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
type Create func(ctx context.Context, perfectBalancedCorpus bool) error

// MakeCreate creates a new Create function
func MakeCreate(selectByCategorizations categorized.SelectByCategorizations, selectTweetByID tweets.SelectByID, selectTweetQuoteByID quotes.SelectByID, deleteAllCorpusRows DeleteAll, insertCorpusRow Insert) Create {
	var categorizations = []string{categorized.VerdictPositive, categorized.VerdictNegative}

	return func(ctx context.Context, perfectBalancedCorpus bool) error {
		categorizedTweets, err := selectByCategorizations(ctx, categorizations)
		if err != nil {
			log.Error(ctx, err.Error())
			return FailedToRetrieveCategorizedTweets
		}

		tweetsBySearchCriteria := make(map[int]map[string][]DTO)
		for _, categorizedTweet := range categorizedTweets {
			tweetData, err := selectTweetByID(ctx, categorizedTweet.TweetID)
			if err != nil {
				log.Error(ctx, err.Error())
				continue
			}

			tweetsByCategorization := tweetsBySearchCriteria[tweetData.SearchCriteriaID]
			if tweetsByCategorization == nil {
				tweetsByCategorization = make(map[string][]DTO)
			}

			row := DTO{
				TweetAuthor:    tweetData.Author,
				TweetAvatar:    tweetData.Avatar,
				TweetText:      tweetData.TextContent,
				TweetImages:    tweetData.Images,
				IsTweetAReply:  tweetData.IsAReply,
				Categorization: categorizedTweet.Categorization,
			}

			if tweetData.QuoteID != nil {
				tweetQuoteData, err := selectTweetQuoteByID(ctx, *tweetData.QuoteID)
				if err != nil {
					log.Error(ctx, err.Error())
				} else {
					row.QuoteAuthor = &tweetQuoteData.Author
					row.QuoteAvatar = tweetQuoteData.Avatar
					row.QuoteText = tweetQuoteData.TextContent
					row.QuoteImages = tweetQuoteData.Images
					row.IsQuoteAReply = &tweetQuoteData.IsAReply
				}
			}

			rows := tweetsByCategorization[row.Categorization]
			if rows == nil {
				rows = make([]DTO, 0, len(tweetsByCategorization)/2)
			}

			rows = append(rows, row)

			tweetsByCategorization[row.Categorization] = rows
			tweetsBySearchCriteria[tweetData.SearchCriteriaID] = tweetsByCategorization
		}

		err = deleteAllCorpusRows(ctx)
		if err != nil {
			log.Error(ctx, err.Error())
			return FailedToCleanUpCorpusTable
		}

		corpusRows := make([]DTO, 0, len(tweetsBySearchCriteria))
		for _, searchCriteria := range tweetsBySearchCriteria {
			categorizedNegative := searchCriteria[categorized.VerdictNegative]
			categorizedPositive := searchCriteria[categorized.VerdictPositive]

			lenNegative := len(categorizedNegative)
			lenPositive := len(categorizedPositive)

			if perfectBalancedCorpus {
				lenNegative = min(lenNegative, lenPositive)
				lenPositive = lenNegative
			}

			corpusRows = append(corpusRows, categorizedNegative[:lenNegative]...)
			corpusRows = append(corpusRows, categorizedPositive[:lenPositive]...)
		}

		var inserted int
		for _, row := range corpusRows {
			_, err := insertCorpusRow(ctx, row)
			if err != nil {
				log.Error(ctx, err.Error())
				continue
			}

			inserted++
		}

		log.Info(ctx, fmt.Sprintf("Inserted %d/%d rows into the corpus table\n", inserted, len(corpusRows)))

		return nil
	}
}
