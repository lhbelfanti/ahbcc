package quotes_test

import (
	"testing"

	"github.com/jackc/pgx/v5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"ahbcc/cmd/api/tweets/quotes"
	"ahbcc/internal/database"
)

func TestMakeInsertSingle_success(t *testing.T) {
	mockPostgresConnection := new(database.MockPostgresConnection)
	mockPgxRow := new(database.MockPgxRow)
	mockPgxRow.On("Scan", mock.Anything).Return(nil).Run(
		func(args mock.Arguments) {
			dest := args.Get(0).([]interface{})
			ptr, ok := dest[0].(*int)
			if !ok {
				t.Errorf("Incorrect data type %T", dest[0])
			}
			*ptr = 1
		},
	)
	mockPostgresConnection.On("QueryRow", mock.Anything, mock.Anything, mock.Anything).Return(mockPgxRow)
	mockQuoteDTO := quotes.MockQuoteDTO()

	insertSingleQuote := quotes.MakeInsertSingle(mockPostgresConnection)

	want := 1
	got, err := insertSingleQuote(mockQuoteDTO)

	assert.Nil(t, err)
	assert.Equal(t, want, got)
	mockPostgresConnection.AssertExpectations(t)
	mockPgxRow.AssertExpectations(t)
}

func TestMakeInsertSingle_failsWhenInsertOperationThrowsError(t *testing.T) {
	mockPostgresConnection := new(database.MockPostgresConnection)
	mockPgxRow := new(database.MockPgxRow)
	mockPgxRow.On("Scan", mock.Anything).Return(pgx.ErrNoRows)
	mockPostgresConnection.On("QueryRow", mock.Anything, mock.Anything, mock.Anything).Return(mockPgxRow)
	mockQuoteDTO := quotes.MockQuoteDTO()

	insertSingleQuote := quotes.MakeInsertSingle(mockPostgresConnection)

	want := quotes.FailedToInsertQuote
	_, got := insertSingleQuote(mockQuoteDTO)

	assert.Equal(t, want, got)
	mockPostgresConnection.AssertExpectations(t)
	mockPgxRow.AssertExpectations(t)
}
