package categorized

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/jackc/pgx/v5"

	"github.com/lhbelfanti/corpus-creator/internal/database"
	"github.com/lhbelfanti/corpus-creator/internal/log"
)

type (
	// SelectAllByUserID returns a struct with all the analyzed tweets divided by year and month
	SelectAllByUserID func(ctx context.Context, userID int) ([]AnalyzedTweetsDTO, error)

	// SelectByUserIDTweetIDAndSearchCriteriaID returns a categorized tweet DAO by user ID, tweet ID and search criteria ID
	SelectByUserIDTweetIDAndSearchCriteriaID func(ctx context.Context, userID, tweetID, searchCriteriaID int) (DAO, error)

	// SelectByCategorizations returns all the categorized tweets seeking by any of the specified categorizations passed by parameter
	SelectByCategorizations func(ctx context.Context, categorizations []string) ([]DAO, error)
)

// MakeSelectAllByUserID creates a new SelectAllByUserID
func MakeSelectAllByUserID(db database.Connection, collectRows database.CollectRows[AnalyzedTweetsDTO]) SelectAllByUserID {
	const query string = `
		SELECT search_criteria_id, tweet_year, tweet_month, COUNT(*) AS analyzed_tweets
		FROM categorized_tweets
		WHERE user_id = $1
		GROUP BY search_criteria_id, tweet_year, tweet_month
		ORDER BY search_criteria_id, tweet_year, tweet_month;
	`

	return func(ctx context.Context, userID int) ([]AnalyzedTweetsDTO, error) {
		rows, err := db.Query(ctx, query, userID)
		if err != nil {
			log.Error(ctx, err.Error())
			return nil, FailedToExecuteSelectAllCategorizedTweetsByUserID
		}

		analyzedTweets, err := collectRows(rows)
		if err != nil {
			log.Error(ctx, err.Error())
			return nil, FailedToExecuteCollectRowsInSelectAllCategorizedTweetsByUserID
		}

		return analyzedTweets, nil
	}
}

// MakeSelectByUserIDTweetIDAndSearchCriteriaID creates a new SelectByUserIDTweetIDAndSearchCriteriaID
func MakeSelectByUserIDTweetIDAndSearchCriteriaID(db database.Connection) SelectByUserIDTweetIDAndSearchCriteriaID {
	const query string = `SELECT id, search_criteria_id, tweet_id, tweet_year, tweet_month, user_id, categorization
						  FROM categorized_tweets
						  WHERE search_criteria_id = $1 AND tweet_id = $2 AND user_id = $3;`

	return func(ctx context.Context, userID, tweetID, searchCriteriaID int) (DAO, error) {
		var categorizedTweet DAO
		err := db.QueryRow(ctx, query, searchCriteriaID, tweetID, userID).Scan(
			&categorizedTweet.ID,
			&categorizedTweet.SearchCriteriaID,
			&categorizedTweet.TweetID,
			&categorizedTweet.TweetYear,
			&categorizedTweet.TweetMonth,
			&categorizedTweet.UserID,
			&categorizedTweet.Categorization,
		)
		if errors.Is(err, pgx.ErrNoRows) {
			log.Error(ctx, err.Error())
			return DAO{}, NoCategorizedTweetFound
		} else if err != nil {
			log.Error(ctx, err.Error())
			return DAO{}, FailedExecuteQueryToRetrieveCategorizedTweetData
		}

		return categorizedTweet, nil
	}
}

// MakeSelectByCategorizations creates a new SelectByCategorizations function
func MakeSelectByCategorizations(db database.Connection, collectRows database.CollectRows[DAO]) SelectByCategorizations {
	const query string = `SELECT id, search_criteria_id, tweet_id, tweet_year, tweet_month, user_id, categorization
						  FROM categorized_tweets
						  WHERE categorization IN (%s)`

	return func(ctx context.Context, categorizations []string) ([]DAO, error) {
		placeholders := make([]string, len(categorizations))
		values := make([]any, len(categorizations))
		for i, status := range categorizations {
			placeholders[i] = fmt.Sprintf("$%d", i+1)
			values[i] = status
		}

		queryToExecute := fmt.Sprintf(query, strings.Join(placeholders, ","))

		rows, err := db.Query(ctx, queryToExecute, values...)
		if err != nil {
			log.Error(ctx, err.Error())
			return nil, FailedToExecuteSelectByCategorizations
		}

		categorizedTweets, err := collectRows(rows)
		if err != nil {
			log.Error(ctx, err.Error())
			return nil, FailedToExecuteCollectRowsInSelectByCategorizations
		}

		return categorizedTweets, nil
	}
}
