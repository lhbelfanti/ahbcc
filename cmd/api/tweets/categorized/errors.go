package categorized

import "errors"

var (
	FailedToExecuteSelectAllCategorizedTweetsByUserID              = errors.New("failed to execute select all categorized tweets by user id")
	FailedToExecuteCollectRowsInSelectAllCategorizedTweetsByUserID = errors.New("failed to execute collect rows in select all categorized tweets by user id")
	AuthorizationTokenIsRequired                                   = errors.New("authorization token is required")
	InvalidTweetID                                                 = errors.New("invalid tweet id")
	InvalidCategorization                                          = errors.New("invalid categorization")
	FailedToRetrieveUserID                                         = errors.New("failed to retrieve user id")
	FailedToRetrieveTweetByID                                      = errors.New("failed to retrieve tweet by id")
	FailedToInsertSingleCategorizedTweet                           = errors.New("failed to insert single categorized tweet")
	FailedToExecuteInsertCategorizedTweet                          = errors.New("failed to execute insert categorized tweet")
	NoCategorizedTweetFound                                        = errors.New("no categorized tweet found")
	FailedExecuteQueryToRetrieveCategorizedTweetData               = errors.New("failed to execute query to retrieve categorized tweet data")
	FailedToCheckIfTheTweetWasAlreadyCategorized                   = errors.New("failed to check if the tweet was already categorized")
	TweetAlreadyCategorized                                        = errors.New("tweet already categorized")
	FailedToExecuteSelectByCategorizations                         = errors.New("failed to execute select by categorizations")
	FailedToExecuteCollectRowsInSelectByCategorizations            = errors.New("failed to execute collect rows in select by categorizations")
)

const (
	AuthorizationTokenRequired                  string = "Authorization token is required"
	InvalidURLParameter                         string = "Invalid url parameter"
	InvalidRequestBody                          string = "Invalid request body"
	FailedToInsertCategorizedTweet              string = "Failed to insert categorized tweet"
	FailedToCategorizeAnAlreadyCategorizedTweet string = "Failed to categorize an already categorized tweet"
)
