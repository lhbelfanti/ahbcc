package summary

import "errors"

var (
	FailedToInsertExecutionSummary                              = errors.New("failed to insert execution summary")
	FailedToRetrieveMonthlyTweetsCountsByYear                   = errors.New("failed to retrieve monthly tweet count by year")
	FailedToExecuteCollectRowsInSelectMonthlyTweetsCountsByYear = errors.New("failed to execute collect rows in select monthly tweet count by year")
	FailedToRetrieveExecutionsSummary                           = errors.New("failed to retrieve executions summary")
	FailedToExecuteCollectRowsInSelectAll                       = errors.New("failed to execute collect rows in select all")
	FailedToDeleteAllSearchCriteriaExecutionsSummary            = errors.New("failed to delete all search criteria executions summary")
)
