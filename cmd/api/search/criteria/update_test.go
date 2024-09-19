package criteria_test

import (
	"context"
	"errors"
	"testing"

	"github.com/jackc/pgx/v5/pgconn"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"ahbcc/cmd/api/search/criteria"
	"ahbcc/internal/database"
)

func TestUpdateExecution_success(t *testing.T) {
	mockPostgresConnection := new(database.MockPostgresConnection)
	mockPostgresConnection.On("Exec", mock.Anything, mock.Anything, mock.Anything).Return(pgconn.CommandTag{}, nil)

	updateExecution := criteria.MakeUpdateExecution(mockPostgresConnection)

	got := updateExecution(context.Background(), 1, criteria.DoneStatus)

	assert.Nil(t, got)
	mockPostgresConnection.AssertExpectations(t)
}

func TestUpdateExecution_failsWhenUpdateOperationThrowsError(t *testing.T) {
	mockPostgresConnection := new(database.MockPostgresConnection)
	mockPostgresConnection.On("Exec", mock.Anything, mock.Anything, mock.Anything).Return(pgconn.CommandTag{}, errors.New("failed to update execution"))

	updateExecution := criteria.MakeUpdateExecution(mockPostgresConnection)

	want := criteria.FailedToUpdateSearchCriteriaExecution
	got := updateExecution(context.Background(), 1, criteria.DoneStatus)

	assert.Equal(t, want, got)
	mockPostgresConnection.AssertExpectations(t)
}
