package rules_test

import (
	"context"
	"errors"
	"testing"

	"github.com/jackc/pgx/v5/pgconn"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"ahbcc/cmd/api/scrubber/rules"
	"ahbcc/internal/database"
)

func TestInsert_success(t *testing.T) {
	mockPostgresConnection := new(database.MockPostgresConnection)
	mockPostgresConnection.On("Exec", mock.Anything, mock.Anything, mock.Anything).Return(pgconn.CommandTag{}, nil)
	mockRulesDTO := rules.MockRulesDTO()

	insertRules := rules.MakeInsert(mockPostgresConnection)

	got := insertRules(context.Background(), mockRulesDTO)

	assert.Nil(t, got)
	mockPostgresConnection.AssertExpectations(t)
}

func TestInsert_failsWhenInsertOperationThrowsError(t *testing.T) {
	mockPostgresConnection := new(database.MockPostgresConnection)
	mockPostgresConnection.On("Exec", mock.Anything, mock.Anything, mock.Anything).Return(pgconn.CommandTag{}, errors.New("failed to insert rules"))
	mockRulesDTO := rules.MockRulesDTO()

	insertRules := rules.MakeInsert(mockPostgresConnection)

	want := rules.FailedToInsertRules
	got := insertRules(context.Background(), mockRulesDTO)

	assert.Equal(t, want, got)
	mockPostgresConnection.AssertExpectations(t)
}
