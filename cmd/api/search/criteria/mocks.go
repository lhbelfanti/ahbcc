package criteria

import (
	"context"
	"time"
)

// MockSelectByID mocks SelectByID function
func MockSelectByID(dao DAO, err error) SelectByID {
	return func(ctx context.Context, id int) (DAO, error) {
		return dao, err
	}
}

// MockEnqueue mocks Enqueue function
func MockEnqueue(err error) Enqueue {
	return func(ctx context.Context, criteriaID int, forced bool) error {
		return err
	}
}

// MockResume mocks Resume function
func MockResume(err error) Resume {
	return func(ctx context.Context, criteriaID int) error {
		return err
	}
}

// MockInit mocks Init function
func MockInit(err error) Init {
	return func(ctx context.Context) error {
		return err
	}
}

// MockCriteriaDAO mocks a criteria.DAO
func MockCriteriaDAO() DAO {
	return DAO{
		ID:               1,
		Name:             "Example",
		AllOfTheseWords:  []string{"word1", "word2"},
		ThisExactPhrase:  "exact phrase",
		AnyOfTheseWords:  []string{"any1", "any2"},
		NoneOfTheseWords: []string{"none1", "none2"},
		TheseHashtags:    []string{"#hashtag1", "#hashtag2"},
		Language:         "es",
		Since:            time.Date(2006, time.January, 1, 0, 0, 0, 0, time.Local),
		Until:            time.Date(2024, time.January, 1, 0, 0, 0, 0, time.Local),
	}
}

// MockScanCriteriaDAOValues mocks the properties of DAO to be used in the Scan function
func MockScanCriteriaDAOValues(dao DAO) []any {
	return []any{
		dao.ID,
		dao.Name,
		dao.AllOfTheseWords,
		dao.ThisExactPhrase,
		dao.AnyOfTheseWords,
		dao.NoneOfTheseWords,
		dao.TheseHashtags,
		dao.Language,
		dao.Since,
		dao.Until,
	}
}

// MockCriteriaDAOSlice mocks a []criteria.DAO
func MockCriteriaDAOSlice() []DAO {
	return []DAO{
		{
			ID:               1,
			Name:             "Example",
			AllOfTheseWords:  []string{"word1", "word2"},
			ThisExactPhrase:  "exact phrase",
			AnyOfTheseWords:  []string{"any1", "any2"},
			NoneOfTheseWords: []string{"none1", "none2"},
			TheseHashtags:    []string{"#hashtag1", "#hashtag2"},
			Language:         "es",
			Since:            time.Date(2006, time.January, 1, 0, 0, 0, 0, time.Local),
			Until:            time.Date(2024, time.January, 1, 0, 0, 0, 0, time.Local),
		},
		{
			ID:               2,
			Name:             "Example",
			AllOfTheseWords:  []string{"word1", "word2"},
			ThisExactPhrase:  "exact phrase",
			AnyOfTheseWords:  []string{"any1", "any2"},
			NoneOfTheseWords: []string{"none1", "none2"},
			TheseHashtags:    []string{"#hashtag1", "#hashtag2"},
			Language:         "es",
			Since:            time.Date(2006, time.January, 1, 0, 0, 0, 0, time.Local),
			Until:            time.Date(2024, time.January, 1, 0, 0, 0, 0, time.Local),
		},
	}
}
