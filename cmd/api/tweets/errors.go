package tweets

import "errors"

const (
	FailedToInsertTweetsIntoDatabase = "Failed to insert tweets into database"
	InvalidRequestBody               = "Invalid request body"
)

var (
	FailedToInsertTweets                                        = errors.New("failed to insert tweets")
	MissingTweetID                                              = errors.New("missing tweet ID")
	MissingTweetSearchCriteriaID                                = errors.New("missing tweet search criteria ID")
	FailedToRetrieveMonthlyTweetsCountsByYear                   = errors.New("failed to retrieve monthly tweet count by year")
	FailedToExecuteCollectRowsInSelectMonthlyTweetsCountsByYear = errors.New("failed to execute select monthly tweet count by year")
)
