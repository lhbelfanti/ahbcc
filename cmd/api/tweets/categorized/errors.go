package categorized

import "errors"

var (
	FailedToExecuteSelectAllCategorizedTweetsByUserID              = errors.New("failed to execute select all categorized tweets by user id")
	FailedToExecuteCollectRowsInSelectAllCategorizedTweetsByUserID = errors.New("failed to execute collect rows in select all categorized tweets by user id")
	AuthorizationTokenIsRequired                                   = errors.New("authorization token is required")
	InvalidSearchCriteriaID                                        = errors.New("invalid search criteria id")
	InvalidTweetID                                                 = errors.New("invalid tweet id")
	InvalidCategorization                                          = errors.New("invalid categorization")
	FailedToRetrieveUserID                                         = errors.New("failed to retrieve user id")
	FailedToInsertSingleCategorizedTweet                           = errors.New("failed to insert single categorized tweet")
	FailedToExecuteInsertCategorizedTweet                          = errors.New("failed to execute insert categorized tweet")
)

const (
	AuthorizationTokenRequired     string = "Authorization token is required"
	InvalidRequestBody             string = "Invalid request body"
	FailedToInsertCategorizedTweet string = "Failed to insert categorized tweet"
)
