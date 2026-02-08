package quotes

import (
	"context"
	"fmt"
	"strings"

	"github.com/lhbelfanti/corpus-creator/internal/database"
	"github.com/lhbelfanti/corpus-creator/internal/log"
)

// DeleteOrphans when the tweet insertion fails, and that tweets contains a quote,
// the quote would be inserted anyway in the tweets_quotes table. Those quotes are orphan quotes,
// because they don't have a tweet referencing them, so they need to be removed after the tweets
// insertion finished. This function covers this case.
type DeleteOrphans func(ctx context.Context, ids []int) error

// MakeDeleteOrphans creates a new DeleteOrphans
func MakeDeleteOrphans(db database.Connection) DeleteOrphans {
	const query string = `
			DELETE FROM tweets_quotes
			WHERE id IN (%s)
				AND id NOT IN (
					SELECT quote_id
					FROM tweets
					WHERE quote_id IS NOT NULL
				);
	`

	return func(ctx context.Context, ids []int) error {
		placeholders := make([]string, len(ids))
		values := make([]any, len(ids))
		for i, id := range ids {
			placeholders[i] = fmt.Sprintf("$%d", i+1)
			values[i] = id
		}

		queryToExecute := fmt.Sprintf(query, strings.Join(placeholders, ","))

		_, err := db.Exec(ctx, queryToExecute, values...)
		if err != nil {
			log.Error(ctx, err.Error())
			return FailedToDeleteOrphanQuotes
		}

		return nil
	}
}
