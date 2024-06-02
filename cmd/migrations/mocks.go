package migrations

import (
	"context"
	"github.com/jackc/pgx/v5/pgconn"
)

type (
	PgxConnExecFunc func(ctx context.Context, sql string, arguments ...interface{}) (pgconn.CommandTag, error)

	// MockPgxConn mock implementation of the PgConnExecutor
	MockPgxConn struct {
		ExecFunc PgxConnExecFunc
	}
)

func (m *MockPgxConn) Exec(ctx context.Context, sql string, arguments ...interface{}) (pgconn.CommandTag, error) {
	return m.ExecFunc(ctx, sql, arguments...)
}

// MockPgxConnStruct create a new MockPgxConn based on the ExecFunc passed by parameter
func MockPgxConnStruct(pgxConnExecFunc PgxConnExecFunc) *MockPgxConn {
	return &MockPgxConn{
		ExecFunc: pgxConnExecFunc,
	}
}

// MockExecFunc mocks the ExecFunc of MockPgxConn
func MockExecFunc(commandTag string, err error) PgxConnExecFunc {
	return func(ctx context.Context, sql string, arguments ...interface{}) (pgconn.CommandTag, error) {
		return pgconn.NewCommandTag(commandTag), err
	}
}
