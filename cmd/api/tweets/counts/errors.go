package counts

import "errors"

var (
	FailedToInsertTweetCounts                  = errors.New("failed to insert tweets counts")
	FailedToUpdateTotalTweets                  = errors.New("failed to update total_tweets column")
	NoTweetsCountsFoundForTheGivenCriteria     = errors.New("no tweets count found for the given criteria")
	FailedToExecuteQueryToRetrieveTweetsCounts = errors.New("failed to execute query to retrieve tweets counts")
)
