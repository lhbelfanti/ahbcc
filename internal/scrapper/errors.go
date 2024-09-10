package scrapper

import "errors"

var (
	FailedToExecuteSelectCriteriaByID = errors.New("failed to execute select criteria by id")
	FailedToExecuteRequest            = errors.New("request failed")
)
