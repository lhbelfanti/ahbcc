package cleaner

import "errors"

var (
	FailedToRetrieveCleaningRulesByPriority = errors.New("failed to retrieve cleaning rules by priority")
	CannotParseRegex                        = errors.New("cannot parse regex")
)
