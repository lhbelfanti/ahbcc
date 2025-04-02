package counts

import "errors"

var (
	FailedToInsertTweetCounts = errors.New("failed to insert tweets counts")
	FailedToUpdateTotalTweets = errors.New("failed to update total_tweets column")
)
