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
	return DTO{
		RuleType:   ruleType,
		SourceText: sourceText,
		TargetText: targetText,
		Priority:   priority,
	}
}

// MockRulesDTOSlice mocks a slice of rule DTOs
func MockRulesDTOSlice() []DTO {
	replacement := "replacement"

	return []DTO{
		MockRuleDTO(RuleBadWord, "badword", nil, 1),
		MockRuleDTO(RulePerson, "person", nil, 2),
		MockRuleDTO(RuleHashtagDelete, "#delete", nil, 3),
		MockRuleDTO(RuleMentionDelete, "@delete", nil, 4),
		MockRuleDTO(RuleTextDelete, "delete", nil, 5),
		MockRuleDTO(RuleReplacement, "replace", &replacement, 6),
	}
}

// MockRuleDAO mocks a rule DAO
func MockRuleDAO(ruleType string, sourceText string, targetText *string, priority int) DAO {
	return DAO{
		RuleType:   ruleType,
		SourceText: sourceText,
		TargetText: targetText,
		Priority:   priority,
	}
}

// MockRulesDAOSlice mocks a slice of rule DAOs
func MockRulesDAOSlice() []DAO {
	replacement := "replacement"

	return []DAO{
		MockRuleDAO(RuleBadWord, "badword", nil, 1),
		MockRuleDAO(RulePerson, "person", nil, 2),
		MockRuleDAO(RuleHashtagDelete, "#delete", nil, 3),
		MockRuleDAO(RuleMentionDelete, "@delete", nil, 4),
		MockRuleDAO(RuleTextDelete, "delete", nil, 5),
		MockRuleDAO(RuleReplacement, "replace", &replacement, 6),
	}
}
