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

func TestSelectAll_success(t *testing.T) {
	mockPostgresConnection := new(database.MockPostgresConnection)
	mockPgxRows := new(database.MockPgxRows)
	mockPostgresConnection.On("Query", mock.Anything, mock.Anything, mock.Anything).Return(mockPgxRows, nil)
	mockCorpusDAOSlice := []corpus.DAO{corpus.MockDAO()}
	mockCollectRows := database.MockCollectRows[corpus.DAO](mockCorpusDAOSlice, nil)

	selectAll := corpus.MakeSelectAll(mockPostgresConnection, mockCollectRows)

	want := mockCorpusDAOSlice
	got, err := selectAll(context.Background())

	assert.Nil(t, err)
	assert.Equal(t, want, got)
	mockPostgresConnection.AssertExpectations(t)
	mockPgxRows.AssertExpectations(t)
}

func TestSelectAll_failsWhenSelectOperationThrowsError(t *testing.T) {
	mockPostgresConnection := new(database.MockPostgresConnection)
	mockPgxRows := new(database.MockPgxRows)
	mockPostgresConnection.On("Query", mock.Anything, mock.Anything, mock.Anything).Return(mockPgxRows, errors.New("failed to select all corpus entries"))
	mockCorpusDAOSlice := []corpus.DAO{corpus.MockDAO()}
	mockCollectRows := database.MockCollectRows[corpus.DAO](mockCorpusDAOSlice, nil)

	selectAll := corpus.MakeSelectAll(mockPostgresConnection, mockCollectRows)

	want := corpus.FailedToRetrieveAllCorpusEntries
	_, got := selectAll(context.Background())

	assert.Equal(t, want, got)
	mockPostgresConnection.AssertExpectations(t)
	mockPgxRows.AssertExpectations(t)
}

func TestSelectAll_failsWhenCollectRowsThrowsError(t *testing.T) {
	mockPostgresConnection := new(database.MockPostgresConnection)
	mockPgxRows := new(database.MockPgxRows)
	mockPostgresConnection.On("Query", mock.Anything, mock.Anything, mock.Anything).Return(mockPgxRows, nil)
	mockCorpusDAOSlice := []corpus.DAO{corpus.MockDAO()}
	mockCollectRows := database.MockCollectRows[corpus.DAO](mockCorpusDAOSlice, errors.New("failed to collect rows"))

	selectAll := corpus.MakeSelectAll(mockPostgresConnection, mockCollectRows)

	want := corpus.FailedToExecuteCollectRowsInSelectAllCorpusEntries
	_, got := selectAll(context.Background())

	assert.Equal(t, want, got)
	mockPostgresConnection.AssertExpectations(t)
	mockPgxRows.AssertExpectations(t)
}
