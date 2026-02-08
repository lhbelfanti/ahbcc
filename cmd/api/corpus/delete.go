package corpus

import (
	"context"

	"github.com/lhbelfanti/corpus-creator/internal/database"
	"github.com/lhbelfanti/corpus-creator/internal/log"
)

// DeleteAll deletes all entries from the corpus table
type DeleteAll func(ctx context.Context) error

// MakeDeleteAll creates a new DeleteAll function
func MakeDeleteAll(db database.Connection) DeleteAll {
	const query string = `DELETE FROM corpus`

	return func(ctx context.Context) error {
		_, err := db.Exec(ctx, query)
		if err != nil {
			log.Error(ctx, err.Error())
			return FailedToDeleteAllCorpusEntries
		}

		return nil
	}
}
