package database

import (
	"context"

	"github.com/stretchr/testify/mock"

	"github.com/jackc/pgx/v5/pgconn"
)

type (
	// MockPostgresConnection mock implementation of the postgres db
	MockPostgresConnection struct {
		mock.Mock
	}
)

func (m *MockPostgresConnection) Exec(ctx context.Context, sql string, arguments ...interface{}) (pgconn.CommandTag, error) {
	args := m.Called(ctx, sql, arguments)
	return args.Get(0).(pgconn.CommandTag), args.Error(1)
}
