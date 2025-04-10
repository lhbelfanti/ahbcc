package executions

import "errors"

const (
	InvalidURLParameter                    string = "Invalid url parameter"
	InvalidRequestBody                     string = "Invalid request body"
	FailedToExecuteInsertCriteriaExecution string = "Failed to execute insert criteria execution"
	FailedToExecuteUpdateCriteriaExecution string = "Failed to execute update criteria execution"
	FailedToExecuteGetExecutionsByStatuses string = "Failed to execute get criteria executions by statuses"
)

var (
	FailedToInsertSearchCriteriaExecution               = errors.New("failed to insert search criteria execution")
	FailedToUpdateSearchCriteriaExecution               = errors.New("failed to update search criteria execution")
	FailedToExecuteSelectSearchCriteriaExecutionByState = errors.New("failed to execute select search criteria execution by state")
	FailedToExecuteCollectRowsInSelectExecutionByState  = errors.New("failed to execute select collect rows in select criteria execution by state")
	FailedToInsertSearchCriteriaExecutionDay            = errors.New("failed to insert search criteria execution day")
	FailedToRetrieveLastDayExecutedDate                 = errors.New("failed to retrieve last day executed date")
	NoExecutionDaysFoundForTheGivenCriteriaID           = errors.New("no execution days found for the given criteria id")
	NoExecutionFoundForTheGivenID                       = errors.New("no execution found for the given id")
	FailedToExecuteQueryToRetrieveExecutionData         = errors.New("failed to execute query to retrieve execution data")
	FailedToBeginTransaction                            = errors.New("failed to begin transaction")
	FailedToExecuteSelectMonthlyTweetsCountsByYear      = errors.New("failed to execute select monthly tweets count by year")
	FailedToExecuteUpsertExecutionSummary               = errors.New("failed to execute upsert execution summary")
	FailedToCommitTransaction                           = errors.New("failed to commit transaction")
)
