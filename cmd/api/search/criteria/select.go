package criteria

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"

	"ahbcc/internal/database"
	"ahbcc/internal/log"
)

type (
	// SelectByID returns a criteria seeking by criteria ID
	SelectByID func(ctx context.Context, id int) (DAO, error)

	// SelectExecutionByState returns all the search criteria executions in certain state
	SelectExecutionByState func(ctx context.Context, state string) ([]ExecutionDAO, error)
)

// MakeSelectByID creates a new SelectByID
func MakeSelectByID(db database.Connection) SelectByID {
	const query string = `
		SELECT id, name, all_of_these_words, this_exact_phrase, any_of_these_words, none_of_these_words, these_hashtags, language, since_date, until_date
		FROM search_criteria
		WHERE id = %d
	`

	return func(ctx context.Context, id int) (DAO, error) {
		queryToExecute := fmt.Sprintf(query, id)

		var criteria DAO
		err := db.QueryRow(ctx, queryToExecute).Scan(&criteria)
		if errors.Is(err, pgx.ErrNoRows) {
			log.Error(ctx, err.Error())
			return DAO{}, FailedToRetrieveCriteriaData
		}

		return criteria, nil
	}
}

// MakeSelectExecutionsByState creates a new SelectExecutionByState
func MakeSelectExecutionsByState(db database.Connection, collectRows database.CollectRows[ExecutionDAO]) SelectExecutionByState {
	const query string = `
		SELECT id, status, search_criteria_id
		FROM search_criteria_executions
		WHERE status = $1
	`

	return func(ctx context.Context, state string) ([]ExecutionDAO, error) {
		rows, err := db.Query(ctx, query, state)
		if err != nil {
			log.Error(ctx, err.Error())
			return nil, FailedToExecuteSelectSearchCriteriaExecutionByState
		}

		executions, err := collectRows(rows)
		if err != nil {
			log.Error(ctx, err.Error())
			return nil, FailedToExecuteSelectCollectRowsInSelectExecutionByState
		}

		return executions, nil
	}
}
