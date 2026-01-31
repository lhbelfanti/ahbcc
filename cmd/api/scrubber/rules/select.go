package rules

import (
	"ahbcc/internal/database"
	"context"
)

// SelectByPriority is a function that retrieves all the rules that matches the given priority
type SelectByPriority func(ctx context.Context, priority int) ([]DTO, error)

// MakeSelectByPriority creates a new SelectByPriority
func MakeSelectByPriority(db database.Connection, collectRows database.CollectRows[DAO]) SelectByPriority {
	const query string = `
			SELECT rule_type, source_text, target_text, priority 
			FROM corpus_cleaning_rules 
			WHERE priority = $1
	`

	return func(ctx context.Context, priority int) ([]DTO, error) {
		
		return nil, nil
	}
}
