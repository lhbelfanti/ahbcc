package criteria

import "ahbcc/internal/scrapper"

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

const (
	PendingStatus    string = "PENDING"
	InProgressStatus string = "IN_PROGRESS"
	DoneStatus       string = "DONE"
)

// toCriteriaDTO converts a criteria.DAO into a scrapper.CriteriaDTO
func (dao DAO) toCriteriaDTO() scrapper.CriteriaDTO {
	return scrapper.CriteriaDTO{
		ID:               dao.ID,
		Name:             dao.Name,
		AllOfTheseWords:  dao.AllOfTheseWords,
		ThisExactPhrase:  dao.ThisExactPhrase,
		AnyOfTheseWords:  dao.AllOfTheseWords,
		NoneOfTheseWords: dao.NoneOfTheseWords,
		TheseHashtags:    dao.TheseHashtags,
		Language:         dao.Language,
		Since:            dao.Since,
		Until:            dao.Until,
	}
}
