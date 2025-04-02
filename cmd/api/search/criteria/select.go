package criteria

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/jackc/pgx/v5"

	"ahbcc/internal/database"
	"ahbcc/internal/log"
)

type (
	// SelectByID returns a criteria seeking by criteria ID
	SelectByID func(ctx context.Context, id int) (DAO, error)

	// SelectAll returns all the criteria of the 'search_criteria' table
	SelectAll func(ctx context.Context) ([]DAO, error)

	// SelectExecutionByID returns an execution seeking by its ID
	SelectExecutionByID func(ctx context.Context, id int) (ExecutionDAO, error)

	// SelectExecutionsByStatuses returns all the search criteria executions in certain state
	SelectExecutionsByStatuses func(ctx context.Context, statuses []string) ([]ExecutionDAO, error)

	// SelectLastDayExecutedByCriteriaID returns the last day executed for the given criteria
	SelectLastDayExecutedByCriteriaID func(ctx context.Context, id int) (ExecutionDayDAO, error)
)

// MakeSelectByID creates a new SelectByID
func MakeSelectByID(db database.Connection) SelectByID {
	const query string = `
		SELECT id, name, all_of_these_words, this_exact_phrase, any_of_these_words, none_of_these_words, these_hashtags, language, since_date, until_date
		FROM search_criteria
		WHERE id = $1
	`

	return func(ctx context.Context, id int) (DAO, error) {
		var criteria DAO
		err := db.QueryRow(ctx, query, id).Scan(
			&criteria.ID,
			&criteria.Name,
			&criteria.AllOfTheseWords,
			&criteria.ThisExactPhrase,
			&criteria.AnyOfTheseWords,
			&criteria.NoneOfTheseWords,
			&criteria.TheseHashtags,
			&criteria.Language,
			&criteria.Since,
			&criteria.Until,
		)
		if errors.Is(err, pgx.ErrNoRows) {
			log.Error(ctx, err.Error())
			return DAO{}, NoCriteriaDataFoundForTheGivenCriteriaID
		} else if err != nil {
			log.Error(ctx, err.Error())
			return DAO{}, FailedExecuteQueryToRetrieveCriteriaData
		}

		return criteria, nil
	}
}

// MakeSelectAll creates a new SelectAll
func MakeSelectAll(db database.Connection, collectRows database.CollectRows[DAO]) SelectAll {
	const query string = `
		SELECT id, name, all_of_these_words, this_exact_phrase, any_of_these_words, none_of_these_words, these_hashtags, language, since_date, until_date
		FROM search_criteria
	`

	return func(ctx context.Context) ([]DAO, error) {
		rows, err := db.Query(ctx, query)
		if err != nil {
			log.Error(ctx, err.Error())
			return nil, FailedToRetrieveAllCriteriaData
		}

		searchCriteria, err := collectRows(rows)
		if err != nil {
			log.Error(ctx, err.Error())
			return nil, FailedToExecuteCollectRowsInSelectAll
		}

		return searchCriteria, nil
	}
}

// MakeSelectExecutionByID creates a new SelectExecutionByID
func MakeSelectExecutionByID(db database.Connection) SelectExecutionByID {
	const query string = `
		SELECT id, status, search_criteria_id
		FROM search_criteria_executions
		WHERE id = $1
	`

	return func(ctx context.Context, id int) (ExecutionDAO, error) {
		var execution ExecutionDAO
		err := db.QueryRow(ctx, query, id).Scan(
			&execution.ID,
			&execution.Status,
			&execution.SearchCriteriaID,
		)
		if errors.Is(err, pgx.ErrNoRows) {
			log.Error(ctx, err.Error())
			return ExecutionDAO{}, NoExecutionFoundForTheGivenID
		} else if err != nil {
			log.Error(ctx, err.Error())
			return ExecutionDAO{}, FailedToExecuteQueryToRetrieveExecutionData
		}

		return execution, nil
	}
}

// MakeSelectExecutionsByStatuses creates a new SelectExecutionsByStatuses
func MakeSelectExecutionsByStatuses(db database.Connection, collectRows database.CollectRows[ExecutionDAO]) SelectExecutionsByStatuses {
	const query string = `
		SELECT id, status, search_criteria_id
		FROM search_criteria_executions
		WHERE status IN (%s)
	`

	return func(ctx context.Context, statuses []string) ([]ExecutionDAO, error) {
		placeholders := make([]string, len(statuses))
		values := make([]any, len(statuses))
		for i, status := range statuses {
			placeholders[i] = fmt.Sprintf("$%d", i+1)
			values[i] = status
		}

		queryToExecute := fmt.Sprintf(query, strings.Join(placeholders, ","))

		rows, err := db.Query(ctx, queryToExecute, values...)
		if err != nil {
			log.Error(ctx, err.Error())
			return nil, FailedToExecuteSelectSearchCriteriaExecutionByState
		}

		executions, err := collectRows(rows)
		if err != nil {
			log.Error(ctx, err.Error())
			return nil, FailedToExecuteCollectRowsInSelectExecutionByState
		}

		return executions, nil
	}
}

// MakeSelectLastDayExecutedByCriteriaID creates a new SelectLastDayExecutedByCriteriaID
func MakeSelectLastDayExecutedByCriteriaID(db database.Connection) SelectLastDayExecutedByCriteriaID {
	const query string = `
		SELECT sced.execution_date, sced.search_criteria_execution_id
		FROM search_criteria_execution_days sced
		JOIN search_criteria_executions sce
		ON sced.search_criteria_execution_id = sce.id
		WHERE sce.search_criteria_id = $1
		ORDER BY sced.execution_date DESC
		LIMIT 1;
	`

	return func(ctx context.Context, criteriaID int) (ExecutionDayDAO, error) {
		var lastExecutionDayExecuted ExecutionDayDAO
		err := db.QueryRow(ctx, query, criteriaID).Scan(
			&lastExecutionDayExecuted.ExecutionDate,
			&lastExecutionDayExecuted.SearchCriteriaExecutionID,
		)
		if errors.Is(err, pgx.ErrNoRows) {
			log.Error(ctx, err.Error())
			return ExecutionDayDAO{}, NoExecutionDaysFoundForTheGivenCriteriaID
		} else if err != nil {
			log.Error(ctx, err.Error())
			return ExecutionDayDAO{}, FailedToRetrieveLastDayExecutedDate
		}

		return lastExecutionDayExecuted, nil
	}
}
