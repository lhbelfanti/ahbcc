package migrations

import (
	"context"

	"github.com/jackc/pgx/v5/pgconn"
)

type (
	// PgxConnExecFunc mock of the Exec function of a PgxConn
	PgxConnExecFunc func(ctx context.Context, sql string, arguments ...interface{}) (pgconn.CommandTag, error)

	// MockPgxConn mock implementation of the PgConnExecutor
	MockPgxConn struct {
		ExecFunc PgxConnExecFunc
	}
)

// MockRun mocks Run function
func MockRun(err error) Run {
	return func(ctx context.Context, migrationsDir string) error {
		return err
	}
}

// MockExecFunc mocks the ExecFunc of MockPgxConn
func MockExecFunc(commandTag string, err error) PgxConnExecFunc {
	return func(ctx context.Context, sql string, arguments ...interface{}) (pgconn.CommandTag, error) {
		return pgconn.NewCommandTag(commandTag), err
	}
}

// MockPgxConnStruct create a new MockPgxConn based on the ExecFunc passed by parameter
func MockPgxConnStruct(pgxConnExecFunc PgxConnExecFunc) *MockPgxConn {
	return &MockPgxConn{
		ExecFunc: pgxConnExecFunc,
	}
}

// Exec mocks the Exec function of a PgxConn
func (m *MockPgxConn) Exec(ctx context.Context, sql string, arguments ...interface{}) (pgconn.CommandTag, error) {
	return m.ExecFunc(ctx, sql, arguments...)
}
