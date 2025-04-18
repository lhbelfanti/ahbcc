package user

import (
	"context"
	"time"
)

// MockExists mocks Exists function
func MockExists(userExists bool, err error) Exists {
	return func(ctx context.Context, username string) (bool, error) {
		return userExists, err
	}
}

// MockSelectByUsername mocks SelectByUsername function
func MockSelectByUsername(userDAO DAO, err error) SelectByUsername {
	return func(ctx context.Context, username string) (DAO, error) {
		return userDAO, err
	}
}

// MockInsert mocks Insert function
func MockInsert(err error) Insert {
	return func(ctx context.Context, user DTO) error {
		return err
	}
}

// MockDTO mocks user DTO
func MockDTO() DTO {
	return DTO{
		Username: "test@test.com",
		Password: "password",
	}
}

// MockDAO mocks user DAO
func MockDAO() DAO {
	return DAO{
		ID:           1,
		Username:     "username",
		PasswordHash: "$2b$12$AEscv2yxile/9dNaaRbIPe2x.9ousol85rtBCPu83cS5F9gXPBZQK\n",
		CreatedAt:    time.Date(2006, time.January, 1, 0, 0, 0, 0, time.Local),
	}
}

// MockScanUserDAOValues mocks the properties of user DAO to be used in the Scan function
func MockScanUserDAOValues(dao DAO) []any {
	return []any{
		dao.ID,
		dao.Username,
		dao.PasswordHash,
		dao.CreatedAt,
	}
}
