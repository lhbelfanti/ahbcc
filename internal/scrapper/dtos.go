package scrapper

type (
	// EnqueueCriteriaMessageDTO is the container that wraps a message to be enqueued
	EnqueueCriteriaMessageDTO struct {
		Message CriteriaDTO `json:"message"`
	}

	// CriteriaDTO represents a search criteria
	CriteriaDTO struct {
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
)
