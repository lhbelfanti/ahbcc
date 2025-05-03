package categorized_test

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"ahbcc/cmd/api/tweets/categorized"
	"ahbcc/internal/database"
)

func TestInsertSingle_success(t *testing.T) {
	mockPostgresConnection := new(database.MockPostgresConnection)
	mockPgxRow := new(database.MockPgxRow)
	database.MockScan(mockPgxRow, []any{1}, t)
	mockPostgresConnection.On("QueryRow", mock.Anything, mock.Anything, mock.Anything).Return(mockPgxRow)
	mockDTO := categorized.MockDTO()

	insertSingle := categorized.MakeInsertSingle(mockPostgresConnection)

	want := 1
	got, err := insertSingle(context.Background(), mockDTO)

	assert.Nil(t, err)
	assert.Equal(t, want, got)
	mockPostgresConnection.AssertExpectations(t)
	mockPgxRow.AssertExpectations(t)
}

func TestInsertSingle_failsWhenScanThrowsError(t *testing.T) {
	mockPostgresConnection := new(database.MockPostgresConnection)
	mockPgxRow := new(database.MockPgxRow)
	mockPgxRow.On("Scan", mock.Anything).Return(errors.New("failed to scan"))
	mockPostgresConnection.On("QueryRow", mock.Anything, mock.Anything, mock.Anything).Return(mockPgxRow)
	mockDTO := categorized.MockDTO()

	insertSingle := categorized.MakeInsertSingle(mockPostgresConnection)

	want := categorized.FailedToExecuteInsertCategorizedTweet
	_, got := insertSingle(context.Background(), mockDTO)

	assert.Equal(t, want, got)
	mockPostgresConnection.AssertExpectations(t)
	mockPgxRow.AssertExpectations(t)
}
