package criteria

// ExecutionDTO represents a search criteria execution to be inserted into the 'search_criteria_executions' table
type ExecutionDTO struct {
	Status           string `json:"status"`
	SearchCriteriaID int    `json:"search_criteria_id"`
}
