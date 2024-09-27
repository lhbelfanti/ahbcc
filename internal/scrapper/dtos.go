package scrapper

type (
	// EnqueueCriteriaMessageDTO is the container that wraps a message to be enqueued
	EnqueueCriteriaMessageDTO struct {
		Message Message `json:"message"`
	}

	// Message is a struct that contains the data of the message
	Message struct {
		Criteria    CriteriaDTO `json:"criteria"`
		ExecutionID int         `json:"execution_id"`
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

// newEnqueueCriteriaMessageDTO creates a new EnqueueCriteriaMessageDTO
func newEnqueueCriteriaMessageDTO(criteria CriteriaDTO, executionID int) EnqueueCriteriaMessageDTO {
	return EnqueueCriteriaMessageDTO{
		Message: Message{
			Criteria:    criteria,
			ExecutionID: executionID,
		},
	}
}
