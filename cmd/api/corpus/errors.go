package corpus

import "errors"

var (
	FailedToInsertCorpusEntry                          = errors.New("failed to insert corpus entry")
	FailedToDeleteAllCorpusEntries                     = errors.New("failed to delete all corpus entries")
	FailedToRetrieveAllCorpusEntries                   = errors.New("failed to retrieve all corpus entries")
	FailedToExecuteCollectRowsInSelectAllCorpusEntries = errors.New("failed to execute collect rows in select all corpus entries")
	FailedToRetrieveCategorizedTweets                  = errors.New("failed to retrieve categorized tweets")
	FailedToCleanUpCorpusTable                         = errors.New("failed to clean up corpus table")
	FailedToCleanTweets                                = errors.New("failed to clean tweets")
	FailedToExecuteSelectAll                           = errors.New("failed to execute select all")
	InvalidExportFormat                                = errors.New("invalid export format")
)

const (
	FailedToCreateCorpus string = "Failed to create corpus"
	FailedToExportCorpus string = "Failed to export corpus"
)
