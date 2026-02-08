package tweets

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"

	"github.com/lhbelfanti/corpus-creator/cmd/api/user/session"
	"github.com/lhbelfanti/corpus-creator/internal/database"
	"github.com/lhbelfanti/corpus-creator/internal/log"
)

type (
	// SelectBySearchCriteriaIDYearAndMonth retrieves the user's uncategorized tweets from a criteria, seeking by year and month.
	// It also limits the number of tweets retrieved to the limit param.
	SelectBySearchCriteriaIDYearAndMonth func(ctx context.Context, searchCriteriaID, year, month, limit int, token string) ([]CustomTweetDTO, error)

	// SelectByID retrieves a tweet DAO by its ID
	SelectByID func(ctx context.Context, id int) (DAO, error)
)

// MakeSelectBySearchCriteriaIDYearAndMonth creates a new SelectBySearchCriteriaIDYearAndMonth
func MakeSelectBySearchCriteriaIDYearAndMonth(db database.Connection, collectRows database.CollectRows[CustomTweetDTO], selectUserIDByToken session.SelectUserIDByToken) SelectBySearchCriteriaIDYearAndMonth {
	const query string = `SELECT t.id, t.status_id, t.author, t.avatar, t.posted_at, t.is_a_reply, t.text_content, t.images, t.quote_id, t.search_criteria_id,
         							q.author, q.avatar, q.posted_at, q.is_a_reply, q.text_content, q.images
						  FROM tweets AS t
						  LEFT JOIN tweets_quotes AS q ON t.quote_id = q.id
						  WHERE t.id NOT IN (SELECT c.tweet_id FROM categorized_tweets AS c WHERE c.user_id = $1)
						    AND t.search_criteria_id = $2`

	return func(ctx context.Context, searchCriteriaID, year, month, limit int, token string) ([]CustomTweetDTO, error) {
		queryToExecute := query

		userID, err := selectUserIDByToken(ctx, token)
		if err != nil {
			log.Error(ctx, err.Error())
			return nil, FailedToRetrieveUserID
		}
		ctx = log.With(ctx, log.Param("user_id", userID))

		args := make([]any, 0, 5)
		currentParamNum := 3

		args = append(args, userID, searchCriteriaID)

		if year != 0 {
			queryToExecute = queryToExecute + fmt.Sprintf(" AND EXTRACT(YEAR FROM t.posted_at) = $%d", currentParamNum)
			currentParamNum++
			args = append(args, year)
			if month != 0 {
				queryToExecute = queryToExecute + fmt.Sprintf(" AND EXTRACT(MONTH FROM t.posted_at) = $%d", currentParamNum)
				currentParamNum++
				args = append(args, month)
			}
		}

		queryToExecute = queryToExecute + fmt.Sprintf(" LIMIT $%d", currentParamNum)
		currentParamNum++
		args = append(args, limit)

		rows, err := db.Query(ctx, queryToExecute, args...)
		if err != nil {
			log.Error(ctx, err.Error())
			return nil, FailedToRetrieveUserUncategorizedTweets
		}

		uncategorizedTweets, err := collectRows(rows)
		if err != nil {
			log.Error(ctx, err.Error())
			return nil, FailedToExecuteCollectRowsInSelectUserUncategorizedTweets
		}

		return uncategorizedTweets, nil
	}
}

// MakeSelectByID creates a new SelectByID
func MakeSelectByID(db database.Connection) SelectByID {
	const query string = `SELECT t.id, t.status_id, t.author, t.avatar, t.posted_at, t.is_a_reply, t.text_content, t.images, t.quote_id, t.search_criteria_id
						  FROM tweets AS t
						  WHERE t.id = $1`

	return func(ctx context.Context, id int) (DAO, error) {
		var tweet DAO
		err := db.QueryRow(ctx, query, id).Scan(
			&tweet.ID,
			&tweet.StatusID,
			&tweet.Author,
			&tweet.Avatar,
			&tweet.PostedAt,
			&tweet.IsAReply,
			&tweet.TextContent,
			&tweet.Images,
			&tweet.QuoteID,
			&tweet.SearchCriteriaID,
		)
		if errors.Is(err, pgx.ErrNoRows) {
			log.Error(ctx, err.Error())
			return DAO{}, NoTweetFoundForTheGivenTweetID
		} else if err != nil {
			log.Error(ctx, err.Error())
			return DAO{}, FailedExecuteQueryToRetrieveTweetData
		}

		return tweet, nil
	}
}
