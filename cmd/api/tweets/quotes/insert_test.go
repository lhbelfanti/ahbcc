package quotes_test

import (
	"context"
	"testing"

	"github.com/jackc/pgx/v5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/lhbelfanti/corpus-creator/cmd/api/tweets/quotes"
	"github.com/lhbelfanti/corpus-creator/internal/database"
)

func TestInsertSingle_success(t *testing.T) {
	mockPostgresConnection := new(database.MockPostgresConnection)
	mockPgxRow := new(database.MockPgxRow)
	database.MockScan(mockPgxRow, []any{1}, t)
	mockPostgresConnection.On("QueryRow", mock.Anything, mock.Anything, mock.Anything).Return(mockPgxRow)
	mockQuoteDTO := quotes.MockQuoteDTO()

	insertSingleQuote := quotes.MakeInsertSingle(mockPostgresConnection)

	want := 1
	got, err := insertSingleQuote(context.Background(), &mockQuoteDTO)

	assert.Nil(t, err)
	assert.Equal(t, want, got)
	mockPostgresConnection.AssertExpectations(t)
	mockPgxRow.AssertExpectations(t)
}

func TestInsertSingle_successEvenWhenTheTimestampParseFails(t *testing.T) {
	mockPostgresConnection := new(database.MockPostgresConnection)
	mockPgxRow := new(database.MockPgxRow)
	database.MockScan(mockPgxRow, []any{1}, t)
	mockPostgresConnection.On("QueryRow", mock.Anything, mock.Anything, mock.Anything).Return(mockPgxRow)
	mockQuoteDTO := quotes.MockQuoteDTO()
	mockQuoteDTO.PostedAt = "wrong"

	insertSingleQuote := quotes.MakeInsertSingle(mockPostgresConnection)

	want := 1
	got, err := insertSingleQuote(context.Background(), &mockQuoteDTO)

	assert.Nil(t, err)
	assert.Equal(t, want, got)
	mockPostgresConnection.AssertExpectations(t)
	mockPgxRow.AssertExpectations(t)
}

func TestInsertSingle_failsWhenQuoteIsNil(t *testing.T) {
	mockPostgresConnection := new(database.MockPostgresConnection)

	insertSingleQuote := quotes.MakeInsertSingle(mockPostgresConnection)

	want := quotes.NothingToInsertWhenQuoteIsNil
	_, got := insertSingleQuote(context.Background(), nil)

	assert.Equal(t, want, got)
	mockPostgresConnection.AssertExpectations(t)
}

func TestInsertSingle_failsWhenInsertOperationThrowsError(t *testing.T) {
	mockPostgresConnection := new(database.MockPostgresConnection)
	mockPgxRow := new(database.MockPgxRow)
	mockPgxRow.On("Scan", mock.Anything).Return(pgx.ErrNoRows)
	mockPostgresConnection.On("QueryRow", mock.Anything, mock.Anything, mock.Anything).Return(mockPgxRow)
	mockQuoteDTO := quotes.MockQuoteDTO()

	insertSingleQuote := quotes.MakeInsertSingle(mockPostgresConnection)

	want := quotes.FailedToInsertQuote
	_, got := insertSingleQuote(context.Background(), &mockQuoteDTO)

	assert.Equal(t, want, got)
	mockPostgresConnection.AssertExpectations(t)
	mockPgxRow.AssertExpectations(t)
}
