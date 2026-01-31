package rules_test

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"ahbcc/cmd/api/scrubber/rules"
	"ahbcc/internal/database"
)

func TestSelectAllByPriority_success(t *testing.T) {
	mockPostgresConnection := new(database.MockPostgresConnection)
	mockPgxRows := new(database.MockPgxRows)
	mockPostgresConnection.On("Query", mock.Anything, mock.Anything, mock.Anything).Return(mockPgxRows, nil)
	mockRulesDAOSlice := rules.MockRulesDAOSlice()
	mockCollectRows := database.MockCollectRows[rules.DAO](mockRulesDAOSlice, nil)

	selectAllByPriority := rules.MakeSelectAllByPriority(mockPostgresConnection, mockCollectRows)

	want := mockRulesDAOSlice
	got, err := selectAllByPriority(context.Background(), 1)

	assert.Nil(t, err)
	assert.Equal(t, want, got)
	mockPostgresConnection.AssertExpectations(t)
	mockPgxRows.AssertExpectations(t)
}

func TestSelectAllByPriority_failsWhenSelectOperationThrowsError(t *testing.T) {
	mockPostgresConnection := new(database.MockPostgresConnection)
	mockPgxRows := new(database.MockPgxRows)
	mockPostgresConnection.On("Query", mock.Anything, mock.Anything, mock.Anything).Return(mockPgxRows, errors.New("failed to select all by priority"))
	mockRulesDAOSlice := rules.MockRulesDAOSlice()
	mockCollectRows := database.MockCollectRows[rules.DAO](mockRulesDAOSlice, nil)

	selectAllByPriority := rules.MakeSelectAllByPriority(mockPostgresConnection, mockCollectRows)

	want := rules.FailedToExecuteSelectAllRulesByPriority
	_, got := selectAllByPriority(context.Background(), 1)

	assert.Equal(t, want, got)
	mockPostgresConnection.AssertExpectations(t)
	mockPgxRows.AssertExpectations(t)
}

func TestSelectAllByPriority_failsWhenCollectRowsThrowsError(t *testing.T) {
	mockPostgresConnection := new(database.MockPostgresConnection)
	mockPgxRows := new(database.MockPgxRows)
	mockPostgresConnection.On("Query", mock.Anything, mock.Anything, mock.Anything).Return(mockPgxRows, nil)
	mockCollectRows := database.MockCollectRows[rules.DAO](nil, errors.New("failed to collect rows"))

	selectAllByPriority := rules.MakeSelectAllByPriority(mockPostgresConnection, mockCollectRows)

	want := rules.FailedToExecuteCollectRowsInSelectAllRulesByPriority
	_, got := selectAllByPriority(context.Background(), 1)

	assert.Equal(t, want, got)
	mockPostgresConnection.AssertExpectations(t)
	mockPgxRows.AssertExpectations(t)
}
