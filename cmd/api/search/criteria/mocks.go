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

// MockSelectAll mocks SelectAll function
func MockSelectAll(daos []DAO, err error) SelectAll {
	return func(ctx context.Context) ([]DAO, error) {
		return daos, err
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

// MockInformation mocks Information function
func MockInformation(informationDTOs InformationDTOs, err error) Information {
	return func(ctx context.Context, token string) (InformationDTOs, error) {
		return informationDTOs, err
	}
}

// MockSummarizedInformation mocks SummarizedInformation function
func MockSummarizedInformation(summarizedInformationDTO SummarizedInformationDTO, err error) SummarizedInformation {
	return func(ctx context.Context, token string, criteriaID int, year int, month int) (SummarizedInformationDTO, error) {
		return summarizedInformationDTO, err
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
			Name:             "Example1",
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
			Name:             "Example2",
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

// MockInformationDTOs mocks a criteria.InformationDTOs
func MockInformationDTOs() InformationDTOs {
	return InformationDTOs{
		{
			Name: "Example1",
			ID:   1,
			Years: YearDataDTOs{
				{
					Year: 2024,
					Months: MonthDataDTOs{
						{Month: 9, AnalyzedTweets: 15, TotalTweets: 350},
					},
				},
				{
					Year: 2025,
					Months: MonthDataDTOs{
						{Month: 1, AnalyzedTweets: 10, TotalTweets: 1000},
						{Month: 2, AnalyzedTweets: 0, TotalTweets: 500},
					},
				},
			},
		},
		{
			Name: "Example2",
			ID:   2,
			Years: YearDataDTOs{
				{
					Year: 2025,
					Months: MonthDataDTOs{
						{Month: 2, AnalyzedTweets: 33, TotalTweets: 333},
						{Month: 5, AnalyzedTweets: 0, TotalTweets: 105},
					},
				},
			},
		},
	}
}

// MockSummarizedInformationDTO mocks a criteria.SummarizedInformationDTO
func MockSummarizedInformationDTO(year, month, analyzed, total int) SummarizedInformationDTO {
	return SummarizedInformationDTO{
		ID:             1,
		Name:           "Example",
		Year:           year,
		Month:          month,
		AnalyzedTweets: analyzed,
		TotalTweets:    total,
	}
}
