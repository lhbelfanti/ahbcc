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

	// SelectExecutionsByStatuses returns all the search criteria executions in certain state
	SelectExecutionsByStatuses func(ctx context.Context, statuses []string) ([]ExecutionDAO, error)
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
		err := db.QueryRow(ctx, query, id).Scan(&criteria)
		if errors.Is(err, pgx.ErrNoRows) {
			log.Error(ctx, err.Error())
			return DAO{}, FailedToRetrieveCriteriaData
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
		defer rows.Close()

		searchCriteria, err := collectRows(rows)
		if err != nil {
			log.Error(ctx, err.Error())
			return nil, FailedToExecuteSelectCollectRowsInSelectAll
		}

		return searchCriteria, nil
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
		for i := range statuses {
			placeholders[i] = fmt.Sprintf("$%d", i+1)
		}

		queryToExecute := fmt.Sprintf(query, strings.Join(placeholders, ","))

		rows, err := db.Query(ctx, queryToExecute, statuses)
		if err != nil {
			log.Error(ctx, err.Error())
			return nil, FailedToExecuteSelectSearchCriteriaExecutionByState
		}
		defer rows.Close()

		executions, err := collectRows(rows)
		if err != nil {
			log.Error(ctx, err.Error())
			return nil, FailedToExecuteSelectCollectRowsInSelectExecutionByState
		}

		return executions, nil
	}
}
