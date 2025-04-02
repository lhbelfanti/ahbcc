package executions

import "time"

type (
	// ExecutionDAO represents a search criteria execution
	ExecutionDAO struct {
		ID               int    `json:"id"`
		Status           string `json:"status"`
		SearchCriteriaID int    `json:"search_criteria_id"`
	}

	// ExecutionDayDAO represents a search criteria execution day
	ExecutionDayDAO struct {
		ID                        int       `json:"id"`
		ExecutionDate             time.Time `json:"execution_date"`
		TweetsQuantity            int       `json:"tweets_quantity"`
		ErrorReason               string    `json:"error_reason"`
		SearchCriteriaExecutionID int       `json:"search_criteria_execution_id"`
	}
)

const (
	PendingStatus    string = "PENDING"
	InProgressStatus string = "IN PROGRESS"
	DoneStatus       string = "DONE"
)
