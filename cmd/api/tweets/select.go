package tweets

import (
	"context"
	"fmt"

	"ahbcc/cmd/api/user/session"
	"ahbcc/internal/database"
	"ahbcc/internal/log"
)

// SelectBySearchCriteriaIDYearAndMonth retrieves the user's uncategorized tweets from a criteria, seeking by year and month.
// It also limits the number of tweets retrieved to the limit param.
type SelectBySearchCriteriaIDYearAndMonth func(ctx context.Context, searchCriteriaID, year, month, limit int, token string) ([]TweetDTO, error)

// MakeSelectBySearchCriteriaIDYearAndMonth creates a new SelectBySearchCriteriaIDYearAndMonth
func MakeSelectBySearchCriteriaIDYearAndMonth(db database.Connection, collectRows database.CollectRows[TweetDTO], selectUserIDByToken session.SelectUserIDByToken) SelectBySearchCriteriaIDYearAndMonth {
	const query string = `SELECT t.id, t.author, t.avatar, t.posted_at, t.is_a_reply, t.text_content, t.images, t.quote_id, t.search_criteria_id
						  FROM tweets as t
						  WHERE t.uuid 
							NOT IN (SELECT c.tweet_id FROM categorized_tweets as c WHERE c.user_id = $1)
						    AND t.search_criteria_id = $2`

	return func(ctx context.Context, searchCriteriaID, year, month, limit int, token string) ([]TweetDTO, error) {
		var queryToExecute string

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
			queryToExecute = query + fmt.Sprintf(" AND year = $%d", currentParamNum)
			currentParamNum++
			args = append(args, year)
			if month != 0 {
				queryToExecute = query + fmt.Sprintf(" AND month = $%d", currentParamNum)
				currentParamNum++
				args = append(args, month)
			}
		}

		queryToExecute = query + fmt.Sprintf(" LIMIT $%d", currentParamNum)
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
