package rules

// DTO represents a rule to be inserted into the 'corpus_cleaning_rules' table
type DTO struct {
	RuleType   string  `json:"rule_type"`
	SourceText string  `json:"source_text"`
	TargetText *string `json:"target_text,omitempty"`
	Priority   int     `json:"priority"`
}

const (
	RuleBadWord       string = "BAD_WORD"
	RulePerson        string = "PERSON"
	RuleHashtagDelete string = "HASHTAG_DELETE"
	RuleMentionDelete string = "MENTION_DELETE"
	RuleTextDelete    string = "TEXT_DELETE"
	RuleReplacement   string = "REPLACEMENT"
)

// validRuleTypes is a list of valid rule types.
// When a new rule type is added, it must be added to this list.
var validRuleTypes = []string{
	RuleBadWord,
	RulePerson,
	RuleHashtagDelete,
	RuleMentionDelete,
	RuleTextDelete,
	RuleReplacement,
}
