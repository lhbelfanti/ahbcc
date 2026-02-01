package rules_test

import (
	rules2 "ahbcc/cmd/api/corpus/cleaner/rules"
	"context"
	"errors"
	"testing"

	"github.com/jackc/pgx/v5/pgconn"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"ahbcc/internal/database"
)

func TestInsert_success(t *testing.T) {
	mockPostgresConnection := new(database.MockPostgresConnection)
	mockPostgresConnection.On("Exec", mock.Anything, mock.Anything, mock.Anything).Return(pgconn.CommandTag{}, nil)
	mockRulesDTOSlice := rules2.MockRulesDTOSlice()

	insertRules := rules2.MakeInsert(mockPostgresConnection)

	got := insertRules(context.Background(), mockRulesDTOSlice)

	assert.Nil(t, got)
	mockPostgresConnection.AssertExpectations(t)
}

func TestInsert_failsWhenInsertOperationThrowsError(t *testing.T) {
	mockPostgresConnection := new(database.MockPostgresConnection)
	mockPostgresConnection.On("Exec", mock.Anything, mock.Anything, mock.Anything).Return(pgconn.CommandTag{}, errors.New("failed to insert rules"))
	mockRulesDTO := rules2.MockRulesDTOSlice()

	insertRules := rules2.MakeInsert(mockPostgresConnection)

	want := rules2.FailedToInsertRules
	got := insertRules(context.Background(), mockRulesDTO)

	assert.Equal(t, want, got)
	mockPostgresConnection.AssertExpectations(t)
}
