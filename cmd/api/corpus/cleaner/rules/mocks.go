package rules

import "context"

// MockInsert mocks Insert function
func MockInsert(err error) Insert {
	return func(ctx context.Context, rules []DTO) error {
		return err
	}
}

// MockSelectAllByPriority mocks SelectAllByPriority function
func MockSelectAllByPriority(cleaningRules []DAO, err error) SelectAllByPriority {
	return func(ctx context.Context, priority int) ([]DAO, error) {
		return cleaningRules, err
	}
}

// MockRuleDTO mocks a rule DTO
func MockRuleDTO(ruleType string, sourceText string, targetText *string, priority int) DTO {
	description := "description"

	return DTO{
		RuleType:    ruleType,
		SourceText:  sourceText,
		TargetText:  targetText,
		Priority:    priority,
		Description: &description,
	}
}

// MockRulesDTOSlice mocks a slice of rule DTOs
func MockRulesDTOSlice() []DTO {
	replacement := "replacement"

	return []DTO{
		MockRuleDTO(RuleBadWord, "badword", nil, 1),
		MockRuleDTO(RuleReplacement, "replace", &replacement, 2),
		MockRuleDTO(RuleDelete, "delete", nil, 3),
	}
}

// MockRuleDAO mocks a rule DAO
func MockRuleDAO(ruleType string, sourceText string, targetText *string, priority int) DAO {
	description := "description"

	return DAO{
		RuleType:    ruleType,
		SourceText:  sourceText,
		TargetText:  targetText,
		Priority:    priority,
		Description: &description,
	}
}

// MockRulesDAOSlice mocks a slice of rule DAOs
func MockRulesDAOSlice() []DAO {
	replacement := "replacement"

	return []DAO{
		MockRuleDAO(RuleBadWord, "badword", nil, 1),
		MockRuleDAO(RuleReplacement, "replace", &replacement, 2),
		MockRuleDAO(RuleDelete, "delete", nil, 3),
	}
}
