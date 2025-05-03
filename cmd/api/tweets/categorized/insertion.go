package categorized

import (
	"context"

	"ahbcc/cmd/api/user/session"
	"ahbcc/internal/log"
)

// InsertCategorizedTweet is a service function that validates the user token and inserts a categorized tweet
type InsertCategorizedTweet func(ctx context.Context, token string, dto DTO) (int, error)

// MakeInsertCategorizedTweet creates a new InsertCategorizedTweet service
func MakeInsertCategorizedTweet(selectUserIDByToken session.SelectUserIDByToken, insertSingle InsertSingle) InsertCategorizedTweet {
	return func(ctx context.Context, token string, dto DTO) (int, error) {
		userID, err := selectUserIDByToken(ctx, token)
		if err != nil {
			log.Error(ctx, err.Error())
			return -1, FailedToRetrieveUserID
		}

		dto.UserID = userID

		categorizedTweetID, err := insertSingle(ctx, dto)
		if err != nil {
			log.Error(ctx, err.Error())
			return -1, FailedToInsertSingleCategorizedTweet
		}

		return categorizedTweetID, nil
	}
}
