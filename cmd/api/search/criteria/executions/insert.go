package executions

import (
	"context"
	"fmt"

	"github.com/lhbelfanti/corpus-creator/internal/database"
	"github.com/lhbelfanti/corpus-creator/internal/log"
)

type (
	// InsertExecution inserts a new search criteria execution into 'search_criteria_executions' table
	InsertExecution func(ctx context.Context, searchCriteriaID int, forced bool) (int, error)

	// InsertExecutionDay inserts a new search criteria execution day into 'search_criteria_execution_days' table
	InsertExecutionDay func(ctx context.Context, executionDay ExecutionDayDTO) error
)

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
			VALUES ('PENDING', %d)
			RETURNING id;
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
			return -1, FailedToInsertSearchCriteriaExecution
		}

		return searchCriteriaExecutionID, nil
	}
}

// MakeInsertExecutionDay creates a new InsertExecutionDay
func MakeInsertExecutionDay(db database.Connection) InsertExecutionDay {
	const query string = `
		INSERT INTO search_criteria_execution_days (execution_date, tweets_quantity, error_reason, search_criteria_execution_id)
		VALUES ($1, $2, $3, $4)
	`
	return func(ctx context.Context, executionDay ExecutionDayDTO) error {
		values := make([]any, 0, 4)
		values = append(values, executionDay.ExecutionDate, executionDay.TweetsQuantity)
		if executionDay.ErrorReason != nil {
			values = append(values, &executionDay.ErrorReason)
		} else {
			values = append(values, nil)
		}

		values = append(values, executionDay.SearchCriteriaExecutionID)

		_, err := db.Exec(ctx, query, values...)
		if err != nil {
			log.Error(ctx, err.Error())
			return FailedToInsertSearchCriteriaExecutionDay
		}

		return nil
	}
}
