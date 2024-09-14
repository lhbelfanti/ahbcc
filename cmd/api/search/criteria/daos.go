package criteria

type (
	// DAO represents a search criteria
	DAO struct {
		ID               int      `json:"id"`
		Name             string   `json:"name"`
		AllOfTheseWords  []string `json:"all_of_these_words"`
		ThisExactPhrase  string   `json:"this_exact_phrase"`
		AnyOfTheseWords  []string `json:"any_of_these_words"`
		NoneOfTheseWords []string `json:"none_of_these_words"`
		TheseHashtags    []string `json:"these_hashtags"`
		Language         string   `json:"language"`
		Since            string   `json:"since"`
		Until            string   `json:"until"`
	}

	// ExecutionDAO represents a search criteria execution
	ExecutionDAO struct {
		ID               int    `json:"id"`
		Status           string `json:"status"`
		SearchCriteriaID int    `json:"search_criteria_id"`
	}
)
