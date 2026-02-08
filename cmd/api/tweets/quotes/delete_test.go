package quotes_test

import (
	"context"
	"errors"
	"testing"

	"github.com/jackc/pgx/v5/pgconn"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/lhbelfanti/corpus-creator/cmd/api/tweets/quotes"
	"github.com/lhbelfanti/corpus-creator/internal/database"
)

func TestDeleteOrphans_success(t *testing.T) {
	mockPostgresConnection := new(database.MockPostgresConnection)
	mockPostgresConnection.On("Exec", mock.Anything, mock.Anything, mock.Anything).Return(pgconn.CommandTag{}, nil)
	mockIDs := []int{1, 2, 3}

	deleteOrphansQuotes := quotes.MakeDeleteOrphans(mockPostgresConnection)

	got := deleteOrphansQuotes(context.Background(), mockIDs)

	assert.Nil(t, got)
	mockPostgresConnection.AssertExpectations(t)
}

func TestDeleteOrphans_failsWhenDeleteOperationThrowsError(t *testing.T) {
	mockPostgresConnection := new(database.MockPostgresConnection)
	mockPostgresConnection.On("Exec", mock.Anything, mock.Anything, mock.Anything).Return(pgconn.CommandTag{}, errors.New("failed to delete orphan quotes"))
	mockIDs := []int{1, 2, 3}

	deleteOrphansQuotes := quotes.MakeDeleteOrphans(mockPostgresConnection)

	want := quotes.FailedToDeleteOrphanQuotes
	got := deleteOrphansQuotes(context.Background(), mockIDs)

	assert.Equal(t, want, got)
	mockPostgresConnection.AssertExpectations(t)
}
