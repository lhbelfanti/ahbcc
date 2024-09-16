package criteria

import "context"

// MockSelectByID mocks SelectByID function
func MockSelectByID(dao DAO, err error) SelectByID {
	return func(ctx context.Context, id int) (DAO, error) {
		return dao, err
	}
}

// MockSelectExecutionsByStatuses mocks SelectExecutionsByStatuses function
func MockSelectExecutionsByStatuses(executionsDAO []ExecutionDAO, err error) SelectExecutionsByStatuses {
	return func(ctx context.Context, statuses []string) ([]ExecutionDAO, error) {
		return executionsDAO, err
	}
}

// MockEnqueue mocks Enqueue function
func MockEnqueue(err error) Enqueue {
	return func(ctx context.Context, criteriaID int, forced bool) error {
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
		Since:            "2006-01-01",
		Until:            "2024-01-01",
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
			Since:            "2006-01-01",
			Until:            "2024-01-01",
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
			Since:            "2006-01-01",
			Until:            "2024-01-01",
		},
	}
}

// MockExecutionDayDTO mocks an ExecutionDayDTO
func MockExecutionDayDTO(errorReason *string) ExecutionDayDTO {
	return ExecutionDayDTO{
		ExecutionDate:             "2006-01-01",
		TweetsQuantity:            20,
		ErrorReason:               errorReason,
		SearchCriteriaExecutionID: 5,
	}
}

// MockExecutionsDAO mocks a slice of ExecutionDAO
func MockExecutionsDAO() []ExecutionDAO {
	return []ExecutionDAO{
		{
			ID:               1,
			Status:           "PENDING",
			SearchCriteriaID: 2,
		},
		{
			ID:               2,
			Status:           "IN PROGRESS",
			SearchCriteriaID: 4,
		},
	}
}
