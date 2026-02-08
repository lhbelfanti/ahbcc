package corpus_test

import (
	"context"
	"errors"
	"testing"

	"github.com/jackc/pgx/v5/pgconn"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/lhbelfanti/corpus-creator/cmd/api/corpus"
	"github.com/lhbelfanti/corpus-creator/internal/database"
)

func TestDeleteAll_success(t *testing.T) {
	mockPostgresConnection := new(database.MockPostgresConnection)
	mockPostgresConnection.On("Exec", mock.Anything, mock.Anything, mock.Anything).Return(pgconn.CommandTag{}, nil)

	deleteAll := corpus.MakeDeleteAll(mockPostgresConnection)

	got := deleteAll(context.Background())

	assert.Nil(t, got)
	mockPostgresConnection.AssertExpectations(t)
}

func TestDeleteAll_failsWhenDBExecFails(t *testing.T) {
	mockPostgresConnection := new(database.MockPostgresConnection)
	mockPostgresConnection.On("Exec", mock.Anything, mock.Anything, mock.Anything).Return(pgconn.CommandTag{}, errors.New("failed to delete corpus rows"))

	deleteAll := corpus.MakeDeleteAll(mockPostgresConnection)

	want := corpus.FailedToDeleteAllCorpusEntries
	got := deleteAll(context.Background())

	assert.Equal(t, want, got)
	mockPostgresConnection.AssertExpectations(t)
}
