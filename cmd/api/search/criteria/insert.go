package criteria

import (
	"context"
	"fmt"

	"ahbcc/internal/database"
	"ahbcc/internal/log"
)

// InsertExecution inserts a new search criteria execution into 'search_criteria_executions' table
type InsertExecution func(ctx context.Context, searchCriteriaID int, forced bool) (int, error)

// MakeInsertExecution creates a new InsertExecution
func MakeInsertExecution(db database.Connection) InsertExecution {
	const (
		query string = `
			INSERT INTO search_criteria_executions (status, search_criteria_id)
			SELECT 'PENDING', %[1]d
			WHERE NOT EXISTS (
				SELECT 1
				FROM search_criteria_executions
				WHERE search_criteria_id = %[1]d
				AND status IN ('PENDING', 'IN PROGRESS')
			)
			RETURNING id;
		`

		forcedInsertQuery string = `
			INSERT INTO search_criteria_executions(status, search_criteria_id)
			VALUES %d
		`
	)

	return func(ctx context.Context, searchCriteriaID int, forced bool) (int, error) {
		var queryToExecute string
		if forced {
			queryToExecute = fmt.Sprintf(forcedInsertQuery, searchCriteriaID)
		} else {
			queryToExecute = fmt.Sprintf(query, searchCriteriaID)
		}

		var searchCriteriaExecutionID int
		err := db.QueryRow(ctx, queryToExecute).Scan(&searchCriteriaExecutionID)
		if err != nil {
			log.Error(ctx, err.Error())
			return -1, FailedToInsertSearchExecutionCriteria
		}

		return searchCriteriaExecutionID, nil
	}
}
