package corpus_test

import (
	"context"
	"errors"
	"testing"

	"github.com/jackc/pgx/v5/pgconn"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"ahbcc/cmd/api/corpus"
	"ahbcc/internal/database"
)

func TestInsert_success(t *testing.T) {
	mockCorpusDTO := corpus.MockDTO()
	mockPostgresConnection := new(database.MockPostgresConnection)
	mockPostgresConnection.On("Exec", mock.Anything, mock.Anything, mock.Anything).Return(pgconn.CommandTag{}, nil)

	insertCorpusEntry := corpus.MakeInsert(mockPostgresConnection)

	got := insertCorpusEntry(context.Background(), mockCorpusDTO)

	assert.Nil(t, got)
	mockPostgresConnection.AssertExpectations(t)
}

func TestInsert_failsWhenInsertOperationThrowsError(t *testing.T) {
	mockCorpusDTO := corpus.MockDTO()
	mockPostgresConnection := new(database.MockPostgresConnection)
	mockPostgresConnection.On("Exec", mock.Anything, mock.Anything, mock.Anything).Return(pgconn.CommandTag{}, errors.New("db error"))

	insertCorpusEntry := corpus.MakeInsert(mockPostgresConnection)

	want := corpus.FailedToInsertCorpusEntry
	got := insertCorpusEntry(context.Background(), mockCorpusDTO)

	assert.Equal(t, want, got)
	mockPostgresConnection.AssertExpectations(t)
}
