package quotes

import "errors"

var (
	NothingToInsertWhenQuoteIsNil              = errors.New("nothing to insert when quote is nil")
	FailedToInsertQuote                        = errors.New("failed to insert quote")
	NoTweetQuoteFoundForTheGivenTweetQuoteID   = errors.New("no tweet quote found for the given tweet quote id")
	FailedExecuteQueryToRetrieveTweetQuoteData = errors.New("failed to execute query to retrieve tweet quote data")
	FailedToDeleteOrphanQuotes                 = errors.New("failed to delete orphan quotes")
)
