package categorized

import (
	"context"
	
	"ahbcc/cmd/api/tweets"
	"ahbcc/cmd/api/user/session"
	"ahbcc/internal/log"
)

// InsertCategorizedTweet inserts a categorized tweet
type InsertCategorizedTweet func(ctx context.Context, token string, tweetID int, body InsertSingleBodyDTO) (int, error)

// MakeInsertCategorizedTweet creates a new InsertCategorizedTweet service
func MakeInsertCategorizedTweet(selectUserIDByToken session.SelectUserIDByToken, selectTweetByID tweets.SelectByID, insertSingle InsertSingle) InsertCategorizedTweet {
	return func(ctx context.Context, token string, tweetID int, body InsertSingleBodyDTO) (int, error) {
		userID, err := selectUserIDByToken(ctx, token)
		if err != nil {
			log.Error(ctx, err.Error())
			return -1, FailedToRetrieveUserID
		}

		tweetDAO, err := selectTweetByID(ctx, tweetID)
		if err != nil {
			log.Error(ctx, err.Error())
			return -1, FailedToRetrieveTweetByID
		}

		categorizedTweet := DTO{
			SearchCriteriaID: tweetDAO.SearchCriteriaID,
			TweetID:          tweetDAO.ID,
			TweetYear:        tweetDAO.PostedAt.Year(),
			TweetMonth:       int(tweetDAO.PostedAt.Month()),
			UserID:           userID,
			Categorization:   body.Categorization,
		}

		categorizedTweetID, err := insertSingle(ctx, categorizedTweet)
		if err != nil {
			log.Error(ctx, err.Error())
			return -1, FailedToInsertSingleCategorizedTweet
		}

		return categorizedTweetID, nil
	}
}
