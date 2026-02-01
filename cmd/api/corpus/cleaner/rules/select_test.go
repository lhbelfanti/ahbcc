package rules_test

import (
	rules2 "ahbcc/cmd/api/corpus/cleaner/rules"
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"ahbcc/internal/database"
)

func TestSelectAllByPriority_success(t *testing.T) {
	mockPostgresConnection := new(database.MockPostgresConnection)
	mockPgxRows := new(database.MockPgxRows)
	mockPostgresConnection.On("Query", mock.Anything, mock.Anything, mock.Anything).Return(mockPgxRows, nil)
	mockRulesDAOSlice := rules2.MockRulesDAOSlice()
	mockCollectRows := database.MockCollectRows[rules2.DAO](mockRulesDAOSlice, nil)

	selectAllByPriority := rules2.MakeSelectAllByPriority(mockPostgresConnection, mockCollectRows)

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
	mockRulesDAOSlice := rules2.MockRulesDAOSlice()
	mockCollectRows := database.MockCollectRows[rules2.DAO](mockRulesDAOSlice, nil)

	selectAllByPriority := rules2.MakeSelectAllByPriority(mockPostgresConnection, mockCollectRows)

	want := rules2.FailedToExecuteSelectAllRulesByPriority
	_, got := selectAllByPriority(context.Background(), 1)

	assert.Equal(t, want, got)
	mockPostgresConnection.AssertExpectations(t)
	mockPgxRows.AssertExpectations(t)
}

func TestSelectAllByPriority_failsWhenCollectRowsThrowsError(t *testing.T) {
	mockPostgresConnection := new(database.MockPostgresConnection)
	mockPgxRows := new(database.MockPgxRows)
	mockPostgresConnection.On("Query", mock.Anything, mock.Anything, mock.Anything).Return(mockPgxRows, nil)
	mockCollectRows := database.MockCollectRows[rules2.DAO](nil, errors.New("failed to collect rows"))

	selectAllByPriority := rules2.MakeSelectAllByPriority(mockPostgresConnection, mockCollectRows)

	want := rules2.FailedToExecuteCollectRowsInSelectAllRulesByPriority
	_, got := selectAllByPriority(context.Background(), 1)

	assert.Equal(t, want, got)
	mockPostgresConnection.AssertExpectations(t)
	mockPgxRows.AssertExpectations(t)
}
