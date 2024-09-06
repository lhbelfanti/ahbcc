package criteria

// MockCriteria mocks a criteria.Type
func MockCriteria() Type {
	return Type{
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
