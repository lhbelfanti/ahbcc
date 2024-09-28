package criteria

type (
	// ExecutionDTO represents a search criteria execution to be inserted into the 'search_criteria_executions' table
	ExecutionDTO struct {
		Status           string `json:"status"`
		SearchCriteriaID int    `json:"search_criteria_id"`
	}

	// ExecutionDayDTO represents a search criteria execution day to be inserted into the 'search_criteria_execution_days' table
	ExecutionDayDTO struct {
		ExecutionDate             string  `json:"execution_date"`
		TweetsQuantity            int     `json:"tweets_quantity"`
		ErrorReason               *string `json:"error_reason"`
		SearchCriteriaExecutionID int     `json:"search_criteria_execution_id"`
	}
)
