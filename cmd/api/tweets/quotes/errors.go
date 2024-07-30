package quotes

import "errors"

var (
	NothingToInsertWhenQuoteIsNil = errors.New("nothing to insert when quote is nil")
	FailedToInsertQuote           = errors.New("failed to insert quote")

	FailedToDeleteOrphanQuotes = errors.New("failed to delete orphan quotes")
)
