package rules

import "context"

// MockInsert mocks Insert function
func MockInsert(err error) Insert {
	return func(ctx context.Context, rules []DTO) error {
		return err
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

// MockRulesDTO mocks a slice of rule DTOs
func MockRulesDTO() []DTO {
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
