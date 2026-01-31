package rules

import "errors"

var (
	FailedToInsertRules                                  = errors.New("failed to insert rules")
	FailedToExecuteSelectAllRulesByPriority              = errors.New("failed to execute select rules by priority")
	FailedToExecuteCollectRowsInSelectAllRulesByPriority = errors.New("failed to execute collect rows in select rules by priority")
)
