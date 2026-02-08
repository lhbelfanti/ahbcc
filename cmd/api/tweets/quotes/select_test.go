package quotes_test

import (
	"context"
	"errors"
	"testing"

	"github.com/jackc/pgx/v5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/lhbelfanti/corpus-creator/cmd/api/tweets/quotes"
	"github.com/lhbelfanti/corpus-creator/internal/database"
)

func TestSelectByID_success(t *testing.T) {
	mockPostgresConnection := new(database.MockPostgresConnection)

	mockPgxRow := new(database.MockPgxRow)
	mockTweetQuoteDAO := quotes.MockTweetQuoteDAO()
	mockScanTweetDAOValues := quotes.MockScanTweetQuoteDAOValues(mockTweetQuoteDAO)
	database.MockScan(mockPgxRow, mockScanTweetDAOValues, t)
	mockPostgresConnection.On("QueryRow", mock.Anything, mock.Anything, mock.Anything).Return(mockPgxRow)

	selectByID := quotes.MakeSelectByID(mockPostgresConnection)

	want := mockTweetQuoteDAO
	got, err := selectByID(context.Background(), 1)

	assert.Nil(t, err)
	assert.Equal(t, want, got)
	mockPostgresConnection.AssertExpectations(t)
	mockPgxRow.AssertExpectations(t)
}

func TestSelectByID_failsWhenSelectOperationFails(t *testing.T) {
	tests := []struct {
		err      error
		expected error
	}{
		{err: pgx.ErrNoRows, expected: quotes.NoTweetQuoteFoundForTheGivenTweetQuoteID},
		{err: errors.New("failed to execute select operation"), expected: quotes.FailedExecuteQueryToRetrieveTweetQuoteData},
	}

	for _, tt := range tests {
		mockPostgresConnection := new(database.MockPostgresConnection)
		mockPgxRow := new(database.MockPgxRow)
		mockPgxRow.On("Scan", mock.Anything).Return(tt.err)
		mockPostgresConnection.On("QueryRow", mock.Anything, mock.Anything, mock.Anything).Return(mockPgxRow)

		selectByID := quotes.MakeSelectByID(mockPostgresConnection)

		want := tt.expected
		_, got := selectByID(context.Background(), 1)

		assert.Equal(t, want, got)
	}
}
