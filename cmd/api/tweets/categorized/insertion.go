package categorized

import (
	"context"
	"errors"

	"github.com/lhbelfanti/corpus-creator/cmd/api/tweets"
	"github.com/lhbelfanti/corpus-creator/cmd/api/user/session"
	"github.com/lhbelfanti/corpus-creator/internal/log"
)

// InsertCategorizedTweet inserts a categorized tweet
type InsertCategorizedTweet func(ctx context.Context, token string, tweetID int, body InsertSingleBodyDTO) (int, error)

// MakeInsertCategorizedTweet creates a new InsertCategorizedTweet service
func MakeInsertCategorizedTweet(selectUserIDByToken session.SelectUserIDByToken, selectTweetByID tweets.SelectByID, selectByUserIDTweetIDAndSearchCriteriaID SelectByUserIDTweetIDAndSearchCriteriaID, insertSingle InsertSingle) InsertCategorizedTweet {
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

		_, err = selectByUserIDTweetIDAndSearchCriteriaID(ctx, userID, tweetID, tweetDAO.SearchCriteriaID)
		if err == nil {
			log.Error(ctx, TweetAlreadyCategorized.Error())
			return -1, TweetAlreadyCategorized
		}

		if err != nil && !errors.Is(err, NoCategorizedTweetFound) {
			log.Error(ctx, err.Error())
			return -1, FailedToCheckIfTheTweetWasAlreadyCategorized
		}

		newCategorizedTweet := DTO{
			SearchCriteriaID: tweetDAO.SearchCriteriaID,
			TweetID:          tweetDAO.ID,
			TweetYear:        tweetDAO.PostedAt.Year(),
			TweetMonth:       int(tweetDAO.PostedAt.Month()),
			UserID:           userID,
			Categorization:   body.Categorization,
		}

		categorizedTweetID, err := insertSingle(ctx, newCategorizedTweet)
		if err != nil {
			log.Error(ctx, err.Error())
			return -1, FailedToInsertSingleCategorizedTweet
		}

		return categorizedTweetID, nil
	}
}
