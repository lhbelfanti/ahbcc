package database

import (
	"context"
	"testing"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/stretchr/testify/mock"
)

type (
	// MockPostgresConnection mock implementation of the postgres db
	MockPostgresConnection struct {
		mock.Mock
	}

	// MockPgxRow mock implementation of pgx.Row
	MockPgxRow struct {
		mock.Mock
	}

	// MockPgxRows mock implementation of pgx.Rows
	MockPgxRows struct {
		mock.Mock
	}
)

func (m *MockPostgresConnection) Exec(ctx context.Context, sql string, arguments ...interface{}) (pgconn.CommandTag, error) {
	args := m.Called(ctx, sql, arguments)
	return args.Get(0).(pgconn.CommandTag), args.Error(1)
}

func (m *MockPostgresConnection) Query(ctx context.Context, sql string, arguments ...any) (pgx.Rows, error) {
	args := m.Called(ctx, sql, arguments)
	return args.Get(0).(pgx.Rows), args.Error(1)
}

func (m *MockPostgresConnection) QueryRow(ctx context.Context, sql string, arguments ...any) pgx.Row {
	args := m.Called(ctx, sql, arguments)
	return args.Get(0).(pgx.Row)
}

func (m *MockPgxRow) Scan(dest ...any) error {
	args := m.Called(dest)
	return args.Error(0)
}

func (m *MockPgxRows) Close() {
	m.Called()
}

func (m *MockPgxRows) Err() error {
	args := m.Called()
	return args.Error(0)
}

func (m *MockPgxRows) CommandTag() pgconn.CommandTag {
	args := m.Called()
	return args.Get(0).(pgconn.CommandTag)
}

func (m *MockPgxRows) FieldDescriptions() []pgconn.FieldDescription {
	args := m.Called()
	return args.Get(0).([]pgconn.FieldDescription)
}

func (m *MockPgxRows) Next() bool {
	args := m.Called()
	return args.Bool(0)
}

func (m *MockPgxRows) Scan(dest ...any) error {
	args := m.Called(dest)
	return args.Error(0)
}

func (m *MockPgxRows) Values() ([]any, error) {
	args := m.Called()
	return args.Get(0).([]any), args.Error(1)
}

func (m *MockPgxRows) RawValues() [][]byte {
	args := m.Called()
	return args.Get(0).([][]byte)
}

func (m *MockPgxRows) Conn() *pgx.Conn {
	args := m.Called()
	return args.Get(0).(*pgx.Conn)
}

// MockScan mocks the "Scan" func
func MockScan[T any](mockPgxRow *MockPgxRow, value T, t *testing.T) {
	mockPgxRow.On("Scan", mock.Anything).Return(nil).Run(
		func(args mock.Arguments) {
			dest := args.Get(0).([]interface{})
			ptr, ok := dest[0].(*T)
			if !ok {
				t.Errorf("Incorrect data type %T", dest[0])
			}
			*ptr = value
		},
	)
}

// MockCollectRows mocks CollectRows function
func MockCollectRows[T any](slice []T, err error) CollectRows[T] {
	return func(rows pgx.Rows) ([]T, error) {
		return slice, err
	}
}
