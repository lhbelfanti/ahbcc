package corpus

import (
	"context"

	"github.com/lhbelfanti/corpus-creator/internal/database"
	"github.com/lhbelfanti/corpus-creator/internal/log"
)

// SelectAll retrieves all entries from the corpus table
type SelectAll func(ctx context.Context) ([]DAO, error)

// MakeSelectAll creates a new SelectAll function
func MakeSelectAll(db database.Connection, collectRows database.CollectRows[DAO]) SelectAll {
	const query string = `SELECT id, tweet_author, tweet_avatar, tweet_text, tweet_images, is_tweet_a_reply,
						  quote_author, quote_avatar, quote_text, quote_images, is_quote_a_reply, categorization
				  		  FROM corpus`

	return func(ctx context.Context) ([]DAO, error) {
		rows, err := db.Query(ctx, query)
		if err != nil {
			log.Error(ctx, err.Error())
			return nil, FailedToRetrieveAllCorpusEntries
		}

		corpusEntries, err := collectRows(rows)
		if err != nil {
			log.Error(ctx, err.Error())
			return nil, FailedToExecuteCollectRowsInSelectAllCorpusEntries
		}

		return corpusEntries, nil
	}
}
