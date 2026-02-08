package criteria

import (
	"time"

	"github.com/lhbelfanti/corpus-creator/internal/scrapper"
)

// DAO represents a search criteria
type DAO struct {
	ID               int       `json:"id"`
	Name             string    `json:"name"`
	AllOfTheseWords  []string  `json:"all_of_these_words"`
	ThisExactPhrase  string    `json:"this_exact_phrase"`
	AnyOfTheseWords  []string  `json:"any_of_these_words"`
	NoneOfTheseWords []string  `json:"none_of_these_words"`
	TheseHashtags    []string  `json:"these_hashtags"`
	Language         string    `json:"language"`
	Since            time.Time `json:"since"`
	Until            time.Time `json:"until"`
}

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
		Since:            dao.Since.Format("2006-01-02"),
		Until:            dao.Until.Format("2006-01-02"),
	}
}
