package executions

import (
	"context"

	"github.com/lhbelfanti/corpus-creator/internal/database"
	"github.com/lhbelfanti/corpus-creator/internal/log"
)

// UpdateExecution updates a search criteria execution status
type UpdateExecution func(ctx context.Context, id int, status string) error

// MakeUpdateExecution creates a new UpdateExecution
func MakeUpdateExecution(db database.Connection) UpdateExecution {
	const query string = `
		UPDATE search_criteria_executions
		SET status = $2
		WHERE id = $1;
	`
	return func(ctx context.Context, id int, status string) error {
		_, err := db.Exec(ctx, query, id, status)
		if err != nil {
			log.Error(ctx, err.Error())
			return FailedToUpdateSearchCriteriaExecution
		}

		return nil
	}
}
