package rules

import (
	"context"
	"github.com/lhbelfanti/corpus-creator/internal/database"
	"github.com/lhbelfanti/corpus-creator/internal/log"
)

// SelectAllByPriority is a function that retrieves all the rules that matches the given priority
type SelectAllByPriority func(ctx context.Context, priority int) ([]DAO, error)

// MakeSelectAllByPriority creates a new SelectAllByPriority
func MakeSelectAllByPriority(db database.Connection, collectRows database.CollectRows[DAO]) SelectAllByPriority {
	const query string = `
			SELECT rule_type, source_text, target_text, priority, description
			FROM corpus_cleaning_rules 
			WHERE priority = $1
	`

	return func(ctx context.Context, priority int) ([]DAO, error) {
		rows, err := db.Query(ctx, query, priority)
		if err != nil {
			log.Error(ctx, err.Error())
			return nil, FailedToExecuteSelectAllRulesByPriority
		}

		collectedRows, err := collectRows(rows)
		if err != nil {
			log.Error(ctx, err.Error())
			return nil, FailedToExecuteCollectRowsInSelectAllRulesByPriority
		}

		return collectedRows, nil
	}
}
