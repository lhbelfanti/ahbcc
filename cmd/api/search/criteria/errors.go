package criteria

import "errors"

var (
	NoCriteriaDataFoundForTheGivenCriteriaID          = errors.New("no criteria data found for given criteria id")
	FailedExecuteQueryToRetrieveCriteriaData          = errors.New("failed to execute query to retrieve criteria data")
	FailedToRetrieveAllCriteriaData                   = errors.New("failed to retrieve all criteria data")
	FailedToExecuteCollectRowsInSelectAll             = errors.New("failed to execute select collect rows in select all")
	FailedToExecuteSelectCriteriaByID                 = errors.New("failed to execute select criteria by id")
	FailedToExecuteSelectLastDayExecutedByCriteriaID  = errors.New("failed to execute select last day executed by criteria id")
	FailedToExecuteSelectExecutionsByStatuses         = errors.New("failed to execute select executions by statuses")
	AnExecutionOfThisCriteriaIDIsAlreadyEnqueued      = errors.New("an execution of this criteria is already enqueued")
	FailedToExecuteEnqueueCriteria                    = errors.New("failed to execute enqueue criteria")
	FailedToRetrieveSearchCriteriaExecutionID         = errors.New("failed to retrieve search criteria execution id")
	FailedToInsertSearchCriteriaExecution             = errors.New("failed to insert search criteria execution")
	NoExecutionDaysFoundForTheGivenCriteriaID         = errors.New("no execution days found for the given criteria id")
	FailedToRetrieveUserID                            = errors.New("failed to retrieve user id")
	FailedToRetrieveSearchCriteriaExecutionsSummaries = errors.New("failed to retrieve search criteria executions summaries")
	FailedToRetrieveSearchCriteria                    = errors.New("failed to retrieve search criteria")
	FailedToRetrieveCategorizedTweetsByUserID         = errors.New("failed to retrieve categorized tweets by user id")
	AuthorizationTokenIsRequired                      = errors.New("authorization token is required")
)

const (
	InvalidURLParameter                          string = "Invalid url parameter"
	InvalidQueryParameterFormat                  string = "Invalid query parameter format"
	FailedToEnqueueCriteria                      string = "Failed to execute enqueue criteria"
	ExecutionWithSameCriteriaIDAlreadyEnqueued   string = "An execution with the same criteria id is already enqueued"
	FailedToExecuteInitCriteria                  string = "Failed to execute init criteria"
	FailedToExecuteCriteriaInformation           string = "Failed to execute criteria information"
	FailedToExecuteCriteriaSummarizedInformation string = "Failed to execute criteria summarized information"
	AuthorizationTokenRequired                   string = "Authorization token is required"
)
