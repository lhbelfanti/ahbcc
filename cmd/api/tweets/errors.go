package tweets

import "errors"

var (
	FailedToInsertTweets         = errors.New("failed to insert tweets")
	MissingTweetHash             = errors.New("missing tweet hash")
	MissingTweetSearchCriteriaID = errors.New("missing tweet search criteria ID")
)

const (
	FailedToInsertTweetsIntoDatabase = "Failed to insert tweets into database"
	InvalidRequestBody               = "Invalid request body"
)
