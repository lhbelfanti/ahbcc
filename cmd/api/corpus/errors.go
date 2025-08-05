package corpus

import "errors"

var (
	FailedToInsertCorpusEntry = errors.New("failed to insert corpus entry")
	FailedToDeleteAllCorpusEntries = errors.New("failed to delete all corpus entries")
	FailedToRetrieveAllCorpusEntries = errors.New("failed to retrieve all corpus entries")
	FailedToExecuteCollectRowsInSelectAllCorpusEntries = errors.New("failed to execute collect rows in select all corpus entries")
)
