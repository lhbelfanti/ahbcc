package counts_test

import (
	"context"
	"testing"

	"github.com/jackc/pgx/v5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"ahbcc/cmd/api/tweets/counts"
	"ahbcc/internal/database"
)

func TestInsert_success(t *testing.T) {
	tweetsCountsID := 1
	mockPostgresConnection := new(database.MockPostgresConnection)
	mockPgxRow := new(database.MockPgxRow)
	database.MockScan(mockPgxRow, []any{tweetsCountsID}, t)
	mockPostgresConnection.On("QueryRow", mock.Anything, mock.Anything, mock.Anything).Return(mockPgxRow)
	mockTweetsCountsDAO := counts.MockTweetsCountsDAO(1, 2025, 1, 1000)

	insertTweetCounts := counts.MakeInsert(mockPostgresConnection)

	want := tweetsCountsID
	got, err := insertTweetCounts(context.Background(), mockTweetsCountsDAO)

	assert.Equal(t, want, got)
	assert.Nil(t, err)
	mockPostgresConnection.AssertExpectations(t)
}

func TestInsert_failsWhenInsertOperationThrowsError(t *testing.T) {
	mockPostgresConnection := new(database.MockPostgresConnection)
	mockPgxRow := new(database.MockPgxRow)
	mockPgxRow.On("Scan", mock.Anything).Return(pgx.ErrNoRows)
	mockPostgresConnection.On("QueryRow", mock.Anything, mock.Anything, mock.Anything).Return(mockPgxRow)
	mockTweetsCountsDAO := counts.MockTweetsCountsDAO(1, 2025, 1, 1000)

	insertTweetCounts := counts.MakeInsert(mockPostgresConnection)

	want := counts.FailedToInsertTweetsCounts
	_, got := insertTweetCounts(context.Background(), mockTweetsCountsDAO)

	assert.Equal(t, want, got)
	mockPostgresConnection.AssertExpectations(t)
	mockPgxRow.AssertExpectations(t)
}
