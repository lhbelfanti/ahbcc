package rules

// DTO represents a rule to be inserted into the 'corpus_cleaning_rules' table
type DTO struct {
	RuleType    string  `json:"rule_type"`
	SourceText  string  `json:"source_text"`
	TargetText  *string `json:"target_text,omitempty"`
	Priority    int     `json:"priority"`
	Description string  `json:"description"`
}

const (
	RuleBadWord     string = "BAD_WORD"
	RuleReplacement string = "REPLACE"
	RuleDelete      string = "DELETE"
)

// validRuleTypes is a list of valid rule types.
// When a new rule type is added, it must be added to this list.
var validRuleTypes = []string{
	RuleBadWord,
	RuleReplacement,
	RuleDelete,
}
