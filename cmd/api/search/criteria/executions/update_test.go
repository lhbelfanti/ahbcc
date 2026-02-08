package executions_test

import (
	"context"
	"errors"
	"testing"

	"github.com/jackc/pgx/v5/pgconn"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/lhbelfanti/corpus-creator/cmd/api/search/criteria/executions"
	"github.com/lhbelfanti/corpus-creator/internal/database"
)

func TestUpdateExecution_success(t *testing.T) {
	mockPostgresConnection := new(database.MockPostgresConnection)
	mockPostgresConnection.On("Exec", mock.Anything, mock.Anything, mock.Anything).Return(pgconn.CommandTag{}, nil)

	updateExecution := executions.MakeUpdateExecution(mockPostgresConnection)

	got := updateExecution(context.Background(), 1, executions.DoneStatus)

	assert.Nil(t, got)
	mockPostgresConnection.AssertExpectations(t)
}

func TestUpdateExecution_failsWhenUpdateOperationThrowsError(t *testing.T) {
	mockPostgresConnection := new(database.MockPostgresConnection)
	mockPostgresConnection.On("Exec", mock.Anything, mock.Anything, mock.Anything).Return(pgconn.CommandTag{}, errors.New("failed to update execution"))

	updateExecution := executions.MakeUpdateExecution(mockPostgresConnection)

	want := executions.FailedToUpdateSearchCriteriaExecution
	got := updateExecution(context.Background(), 1, executions.DoneStatus)

	assert.Equal(t, want, got)
	mockPostgresConnection.AssertExpectations(t)
}
