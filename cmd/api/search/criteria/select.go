package criteria

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"

	"ahbcc/internal/database"
	"ahbcc/internal/log"
)

// SelectByID returns a criteria seeking by criteria ID
type SelectByID func(ctx context.Context, id int) (DAO, error)

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
