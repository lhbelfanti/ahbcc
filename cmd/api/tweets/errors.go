package tweets

import "errors"

const (
	InvalidURLParameter              string = "Invalid url parameter"
	InvalidRequestBody               string = "Invalid request body"
	InvalidQueryParameterFormat      string = "Invalid query parameter format"
	AuthorizationTokenRequired       string = "Authorization token is required"
	FailedToInsertTweetsIntoDatabase string = "Failed to insert tweets into database"
	FailedToRetrieveTweets           string = "Failed to insert tweets"
)

var (
	FailedToInsertTweets                                      = errors.New("failed to insert tweets")
	MissingTweetID                                            = errors.New("missing tweet ID")
	MissingTweetSearchCriteriaID                              = errors.New("missing tweet search criteria ID")
	FailedToRetrieveUserUncategorizedTweets                   = errors.New("failed to retrieve user uncategorized tweets")
	FailedToExecuteCollectRowsInSelectUserUncategorizedTweets = errors.New("failed to execute collect rows in select user uncategorized tweets")
	AuthorizationTokenIsRequired                              = errors.New("authorization token is required")
	FailedToRetrieveUserID                                    = errors.New("failed to retrieve user id")
)
