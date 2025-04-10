package summary

import "errors"

var (
	FailedToInsertExecutionSummary                              = errors.New("failed to insert execution summary")
	FailedToUpdateTotalTweets                                   = errors.New("failed to update total_tweets column")
	NoExecutionSummaryFoundForTheGivenCriteria                  = errors.New("no execution summary found for the given criteria")
	FailedToExecuteQueryToRetrieveExecutionsSummary             = errors.New("failed to execute query to retrieve executions summary")
	FailedToRetrieveExecutionSummaryID                          = errors.New("failed to retrieve execution summary ID")
	FailedToRetrieveMonthlyTweetsCountsByYear                   = errors.New("failed to retrieve monthly tweet count by year")
	FailedToExecuteCollectRowsInSelectMonthlyTweetsCountsByYear = errors.New("failed to execute select monthly tweet count by year")
)
