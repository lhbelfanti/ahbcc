package rules

import (
	"context"
	"fmt"
	"strings"

	"ahbcc/internal/database"
	"ahbcc/internal/log"
)

// Insert represents a function that inserts a slice of rules into the 'corpus_cleaning_rules' table
type Insert func(ctx context.Context, rules []DTO) error

// MakeInsert creates a new Insert
func MakeInsert(db database.Connection) Insert {
	const (
		query string = ` 
			INSERT INTO corpus_cleaning_rules(rule_type, source_text, target_text, priority)
			VALUES %s
		    ON CONFLICT (rule_type, source_text, priority) DO NOTHING;
		`

		parameters = 4
	)

	return func(ctx context.Context, rules []DTO) error {
		placeholders := make([]string, 0, len(rules)*parameters)
		values := make([]any, 0, len(rules)*parameters)
		for i, rule := range rules {
			idx := i * parameters
			placeholders = append(placeholders, fmt.Sprintf("($%d::cleaning_rules, $%d, $%d, $%d)", idx+1, idx+2, idx+3, idx+4))
			values = append(values, rule.RuleType, rule.SourceText, rule.TargetText, rule.Priority)
		}

		queryToExecute := fmt.Sprintf(query, strings.Join(placeholders, ","))

		_, err := db.Exec(ctx, queryToExecute, values...)
		if err != nil {
			log.Error(ctx, err.Error())
			return FailedToInsertRules
		}

		return nil
	}
}
