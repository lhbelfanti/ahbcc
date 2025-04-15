package categorized

import "errors"

var (
	FailedToExecuteSelectAllCategorizedTweetsByUserID              = errors.New("failed to execute select all categorized tweets by user id")
	FailedToExecuteCollectRowsInSelectAllCategorizedTweetsByUserID = errors.New("failed to execute collect rows in select all categorized tweets by user id")
)
