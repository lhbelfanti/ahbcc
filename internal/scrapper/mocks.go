package scrapper

import "context"

// MockEnqueueCriteria mocks EnqueueCriteria function
func MockEnqueueCriteria(err error) EnqueueCriteria {
	return func(ctx context.Context, body CriteriaDTO) error {
		return err
	}
}

// MockCriteriaDTO mocks a CriteriaDTO
func MockCriteriaDTO() CriteriaDTO {
	return CriteriaDTO{
		ID:               1,
		Name:             "Example",
		AllOfTheseWords:  []string{"word1", "word2"},
		ThisExactPhrase:  "exact phrase",
		AnyOfTheseWords:  []string{"any1", "any2"},
		NoneOfTheseWords: []string{"none1", "none2"},
		TheseHashtags:    []string{"#hashtag1", "#hashtag2"},
		Language:         "es",
		Since:            "2006-01-01",
		Until:            "2024-01-01",
	}
}
