package rules

// DAO represents a rule from the 'corpus_cleaning_rules' table
type DAO struct {
	RuleType    string  `json:"rule_type"`
	SourceText  string  `json:"source_text"`
	TargetText  *string `json:"target_text,omitempty"`
	Priority    int     `json:"priority"`
	Description string  `json:"description"`
}
