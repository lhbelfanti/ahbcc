package database

import (
	"context"
	"testing"
	"time"

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

	// MockPgxTx mock implementation of pgx.Tx
	MockPgxTx struct {
		mock.Mock
	}

	// MockPgxCollectableRow mock implementation of pgx.CollectableRow
	MockPgxCollectableRow struct {
		mock.Mock
	}
)

// MockPostgresConnection

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

func (m *MockPostgresConnection) Begin(ctx context.Context) (pgx.Tx, error) {
	args := m.Called(ctx)
	return args.Get(0).(pgx.Tx), args.Error(1)
}

// MockPgxRow

func (m *MockPgxRow) Scan(dest ...any) error {
	args := m.Called(dest)
	return args.Error(0)
}

// MockPgxRows

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

// MockPgxTx

func (t *MockPgxTx) Begin(ctx context.Context) (pgx.Tx, error) {
	args := t.Called(ctx)
	return args.Get(0).(pgx.Tx), args.Error(1)
}

func (t *MockPgxTx) Commit(ctx context.Context) error {
	args := t.Called(ctx)
	return args.Error(0)
}

func (t *MockPgxTx) Rollback(ctx context.Context) error {
	args := t.Called(ctx)
	return args.Error(0)
}

func (t *MockPgxTx) CopyFrom(ctx context.Context, tableName pgx.Identifier, columnNames []string, rowSrc pgx.CopyFromSource) (int64, error) {
	args := t.Called(ctx, tableName, columnNames, rowSrc)
	return args.Get(0).(int64), args.Error(1)
}

func (t *MockPgxTx) SendBatch(ctx context.Context, b *pgx.Batch) pgx.BatchResults {
	args := t.Called(ctx, b)
	return args.Get(0).(pgx.BatchResults)
}

func (t *MockPgxTx) LargeObjects() pgx.LargeObjects {
	args := t.Called()
	return args.Get(0).(pgx.LargeObjects)
}

func (t *MockPgxTx) Prepare(ctx context.Context, name, sql string) (*pgconn.StatementDescription, error) {
	args := t.Called(ctx, name, sql)
	return args.Get(0).(*pgconn.StatementDescription), args.Error(1)
}

func (t *MockPgxTx) Exec(ctx context.Context, sql string, arguments ...any) (commandTag pgconn.CommandTag, err error) {
	args := t.Called(ctx, sql, arguments)
	return args.Get(0).(pgconn.CommandTag), args.Error(1)
}

func (t *MockPgxTx) Query(ctx context.Context, sql string, arguments ...any) (pgx.Rows, error) {
	args := t.Called(ctx, sql, arguments)
	return args.Get(0).(pgx.Rows), args.Error(1)
}

func (t *MockPgxTx) QueryRow(ctx context.Context, sql string, arguments ...any) pgx.Row {
	args := t.Called(ctx, sql, arguments)
	return args.Get(0).(pgx.Row)
}

func (t *MockPgxTx) Conn() *pgx.Conn {
	args := t.Called()
	return args.Get(0).(*pgx.Conn)
}

// MockPgxCollectableRow

func (m *MockPgxCollectableRow) Scan(dest ...any) error {
	args := m.Called(dest)
	return args.Error(0)
}

func (m *MockPgxCollectableRow) FieldDescriptions() []pgconn.FieldDescription {
	args := m.Called()
	return args.Get(0).([]pgconn.FieldDescription)
}

func (m *MockPgxCollectableRow) Values() ([]any, error) {
	args := m.Called()
	return args.Get(0).([]any), args.Error(1)
}

func (m *MockPgxCollectableRow) RawValues() [][]byte {
	args := m.Called()
	return args.Get(0).([][]byte)
}

// MockScan mocks the "Scan" func
func MockScan(mockPgxRow *MockPgxRow, values []any, t *testing.T) {
	mockPgxRow.On("Scan", mock.Anything).Return(nil).Run(
		func(args mock.Arguments) {
			dest := args.Get(0).([]interface{})
			if len(dest) != len(values) {
				t.Errorf("Expected %d destination arguments but got %d", len(values), len(dest))
			}
			for i, val := range values {
				switch v := dest[i].(type) {
				case *int:
					*v = val.(int)
				case *bool:
					*v = val.(bool)
				case *string:
					*v = val.(string)
				case *time.Time:
					*v = val.(time.Time)
				case *[]string:
					*v = val.([]string)
				default:
					t.Errorf("Unsupported type %T", v)
				}
			}
		},
	)
}

// MockCollectRows mocks CollectRows function
func MockCollectRows[T any](slice []T, err error) CollectRows[T] {
	return func(rows pgx.Rows) ([]T, error) {
		return slice, err
	}
}

// MockFieldDescriptions mocks field descriptions of a MockPgxCollectableRow
func MockFieldDescriptions(fields []string) []pgconn.FieldDescription {
	descriptions := make([]pgconn.FieldDescription, len(fields))
	for i, name := range fields {
		descriptions[i] = pgconn.FieldDescription{
			Name:                 name,
			TableOID:             0,
			TableAttributeNumber: 0,
			DataTypeOID:          0,
			DataTypeSize:         0,
			TypeModifier:         0,
			Format:               0,
		}
	}
	return descriptions
}

// MockPgxCollectableRowMethods mocks all the methods of a MockPgxCollectableRow
func MockPgxCollectableRowMethods(m *MockPgxCollectableRow, values []any, t *testing.T) {
	m.On("Scan", mock.Anything).Return(nil).Run(func(args mock.Arguments) {
		dest := args.Get(0).([]interface{})
		if len(dest) != len(values) {
			t.Errorf("Expected %d destination arguments but got %d", len(values), len(dest))
			return
		}

		for i, val := range values {
			if val == nil {
				continue // Let the zero value remain
			}

			switch d := dest[i].(type) {
			case *int:
				*d = val.(int)
			case **int:
				*d = val.(*int)
			case *string:
				*d = val.(string)
			case **string:
				*d = val.(*string)
			case *time.Time:
				*d = val.(time.Time)
			case *bool:
				*d = val.(bool)
			case *[]string:
				*d = val.([]string)
			default:
				t.Errorf("Unsupported type %T", d)
			}
		}
	})
}
