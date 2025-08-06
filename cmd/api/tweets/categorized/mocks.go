package categorized

import "context"

// MockSelectAllByUserID mocks a SelectAllByUserID function
func MockSelectAllByUserID(dtos []AnalyzedTweetsDTO, err error) SelectAllByUserID {
	return func(ctx context.Context, userID int) ([]AnalyzedTweetsDTO, error) {
		return dtos, err
	}
}

// MockSelectByUserIDTweetIDAndSearchCriteriaID mocks a SelectByUserIDTweetIDAndSearchCriteriaID function
func MockSelectByUserIDTweetIDAndSearchCriteriaID(dao DAO, err error) SelectByUserIDTweetIDAndSearchCriteriaID {
	return func(ctx context.Context, userID, tweetID, searchCriteriaID int) (DAO, error) {
		return dao, err
	}
}

// MockSelectByCategorizations mocks a SelectByCategorizations function
func MockSelectByCategorizations(daos []DAO, err error) SelectByCategorizations {
	return func(ctx context.Context, categorizations []string) ([]DAO, error) {
		return daos, err
	}
}

// MockInsertSingle mocks an InsertSingle function
func MockInsertSingle(id int, err error) InsertSingle {
	return func(ctx context.Context, dto DTO) (int, error) {
		return id, err
	}
}

// MockInsertCategorizedTweet mocks an InsertCategorizedTweet function
func MockInsertCategorizedTweet(id int, err error) InsertCategorizedTweet {
	return func(ctx context.Context, token string, tweetID int, body InsertSingleBodyDTO) (int, error) {
		return id, err
	}
}

// MockCategorizedTweetsDAO mocks an AnalyzedTweetsDTO
func MockCategorizedTweetsDAO(searchCriteriaID, year, month, analyzed int) AnalyzedTweetsDTO {
	return AnalyzedTweetsDTO{
		SearchCriteriaID: searchCriteriaID,
		Year:             year,
		Month:            month,
		Analyzed:         analyzed,
	}
}

// MockCategorizedTweetsDAOSlice mocks a []categorized.AnalyzedTweetsDTO
func MockCategorizedTweetsDAOSlice() []AnalyzedTweetsDTO {
	return []AnalyzedTweetsDTO{
		MockCategorizedTweetsDAO(1, 2024, 9, 15),
		MockCategorizedTweetsDAO(1, 2025, 1, 10),
		MockCategorizedTweetsDAO(2, 2025, 2, 33),
	}
}

// MockDTO mocks a DTO
func MockDTO() DTO {
	return DTO{
		SearchCriteriaID: 1,
		TweetID:          123456,
		TweetYear:        2024,
		TweetMonth:       5,
		UserID:           789,
		Categorization:   VerdictPositive,
	}
}

// MockInsertSingleBodyDTO mocks an InsertSingleBodyDTO
func MockInsertSingleBodyDTO(verdict string) InsertSingleBodyDTO {
	return InsertSingleBodyDTO{
		Categorization: verdict,
	}
}

// MockCategorizedTweetDAO mocks a DAO
func MockCategorizedTweetDAO() DAO {
	return DAO{
		ID:               1,
		SearchCriteriaID: 2,
		TweetID:          123,
		TweetYear:        2024,
		TweetMonth:       5,
		UserID:           456,
		Categorization:   VerdictPositive,
	}
}

// MockScanCategorizedTweetsDAOValues mocks the properties of DAO to be used in the Scan function
func MockScanCategorizedTweetsDAOValues(dao DAO) []any {
	return []any{
		dao.ID,
		dao.SearchCriteriaID,
		dao.TweetID,
		dao.TweetYear,
		dao.TweetMonth,
		dao.UserID,
		dao.Categorization,
	}
}
