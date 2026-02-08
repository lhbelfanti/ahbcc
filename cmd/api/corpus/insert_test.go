package corpus_test

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/lhbelfanti/corpus-creator/cmd/api/corpus"
	"github.com/lhbelfanti/corpus-creator/internal/database"
)

func TestInsert_success(t *testing.T) {
	mockPostgresConnection := new(database.MockPostgresConnection)
	mockPgxRow := new(database.MockPgxRow)
	database.MockScan(mockPgxRow, []any{1}, t)
	mockPostgresConnection.On("QueryRow", mock.Anything, mock.Anything, mock.Anything).Return(mockPgxRow)
	mockCorpusDTO := corpus.MockDTO()

	insertCorpusEntry := corpus.MakeInsert(mockPostgresConnection)

	want := 1
	got, err := insertCorpusEntry(context.Background(), mockCorpusDTO)

	assert.Nil(t, err)
	assert.Equal(t, want, got)
	mockPostgresConnection.AssertExpectations(t)
	mockPgxRow.AssertExpectations(t)
}

func TestInsert_failsWhenInsertOperationThrowsError(t *testing.T) {
	mockPostgresConnection := new(database.MockPostgresConnection)
	mockPgxRow := new(database.MockPgxRow)
	mockPgxRow.On("Scan", mock.Anything).Return(errors.New("failed to scan"))
	mockPostgresConnection.On("QueryRow", mock.Anything, mock.Anything, mock.Anything).Return(mockPgxRow)
	mockCorpusDTO := corpus.MockDTO()

	insertCorpusEntry := corpus.MakeInsert(mockPostgresConnection)

	want := corpus.FailedToInsertCorpusEntry
	_, got := insertCorpusEntry(context.Background(), mockCorpusDTO)

	assert.Equal(t, want, got)
	mockPostgresConnection.AssertExpectations(t)
	mockPgxRow.AssertExpectations(t)
}
