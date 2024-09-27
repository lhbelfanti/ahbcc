package criteria

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"time"
)

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

// MockSelectLastDayExecutedByCriteriaID mocks SelectLastDayExecutedByCriteriaID function
func MockSelectLastDayExecutedByCriteriaID(lastDayExecuted time.Time, err error) SelectLastDayExecutedByCriteriaID {
	return func(ctx context.Context, id int) (time.Time, error) {
		return lastDayExecuted, err
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

// MockInsertExecution mocks InsertExecution function
func MockInsertExecution(criteriaID int, err error) InsertExecution {
	return func(ctx context.Context, searchCriteriaID int, forced bool) (int, error) {
		return criteriaID, err
	}
}

// MockUpdateExecution mocks UpdateExecution function
func MockUpdateExecution(err error) UpdateExecution {
	return func(ctx context.Context, executionID int, status string) error {
		return err
	}
}

// MockInsertExecutionDay mocks InsertExecutionDay function
func MockInsertExecutionDay(err error) InsertExecutionDay {
	return func(ctx context.Context, executionDay ExecutionDayDTO) error {
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

// MockExecutionsDAO mocks a slice of ExecutionDAO
func MockExecutionsDAO() []ExecutionDAO {
	return []ExecutionDAO{
		{
			ID:               1,
			Status:           PendingStatus,
			SearchCriteriaID: 2,
		},
		{
			ID:               2,
			Status:           InProgressStatus,
			SearchCriteriaID: 4,
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

// MockExecutionDTO mocks an ExecutionDTO
func MockExecutionDTO() ExecutionDTO {
	return ExecutionDTO{
		Status:           "DONE",
		SearchCriteriaID: 0,
	}
}

// MockErrorResponseWriter mocks a ResponseRecorder
type MockErrorResponseWriter struct {
	ResponseRecorder *httptest.ResponseRecorder
}

// Write is the method of MockErrorResponseWriter that mocks the Write behaviour of a ResponseRecorder
func (w *MockErrorResponseWriter) Write(b []byte) (int, error) {
	return 0, errors.New("error while executing ResponseRecorder.Write")
}

// WriteHeader is the method of MockErrorResponseWriter that mocks the WriteHeader behaviour of a ResponseRecorder
func (w *MockErrorResponseWriter) WriteHeader(statusCode int) {
	w.ResponseRecorder.WriteHeader(statusCode)
}

// Header is the method of MockErrorResponseWriter that mocks the Header behaviour of a ResponseRecorder
func (w *MockErrorResponseWriter) Header() http.Header {
	return w.ResponseRecorder.Header()
}
