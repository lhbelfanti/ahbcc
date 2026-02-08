package criteria

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"

	"github.com/lhbelfanti/corpus-creator/internal/database"
	"github.com/lhbelfanti/corpus-creator/internal/log"
)

type (
	// SelectByID returns a criteria seeking by criteria ID
	SelectByID func(ctx context.Context, id int) (DAO, error)

	// SelectAll returns all the criteria of the 'search_criteria' table
	SelectAll func(ctx context.Context) ([]DAO, error)
)

// MakeSelectByID creates a new SelectByID
func MakeSelectByID(db database.Connection) SelectByID {
	const query string = `
		SELECT id, name, all_of_these_words, this_exact_phrase, any_of_these_words, none_of_these_words, these_hashtags, language, since_date, until_date
		FROM search_criteria
		WHERE id = $1;
	`

	return func(ctx context.Context, id int) (DAO, error) {
		var criteria DAO
		err := db.QueryRow(ctx, query, id).Scan(
			&criteria.ID,
			&criteria.Name,
			&criteria.AllOfTheseWords,
			&criteria.ThisExactPhrase,
			&criteria.AnyOfTheseWords,
			&criteria.NoneOfTheseWords,
			&criteria.TheseHashtags,
			&criteria.Language,
			&criteria.Since,
			&criteria.Until,
		)
		if errors.Is(err, pgx.ErrNoRows) {
			log.Error(ctx, err.Error())
			return DAO{}, NoCriteriaDataFoundForTheGivenCriteriaID
		} else if err != nil {
			log.Error(ctx, err.Error())
			return DAO{}, FailedExecuteQueryToRetrieveCriteriaData
		}

		return criteria, nil
	}
}

// MakeSelectAll creates a new SelectAll
func MakeSelectAll(db database.Connection, collectRows database.CollectRows[DAO]) SelectAll {
	const query string = `
		SELECT id, name, all_of_these_words, this_exact_phrase, any_of_these_words, none_of_these_words, these_hashtags, language, since_date, until_date
		FROM search_criteria;
	`

	return func(ctx context.Context) ([]DAO, error) {
		rows, err := db.Query(ctx, query)
		if err != nil {
			log.Error(ctx, err.Error())
			return nil, FailedToRetrieveAllCriteriaData
		}

		searchCriteria, err := collectRows(rows)
		if err != nil {
			log.Error(ctx, err.Error())
			return nil, FailedToExecuteCollectRowsInSelectAll
		}

		return searchCriteria, nil
	}
}
