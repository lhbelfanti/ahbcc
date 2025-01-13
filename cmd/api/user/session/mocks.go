package session

import (
	"context"
	"time"
)

// MockInsertUserSession mocks Insert function
func MockInsertUserSession(err error) Insert {
	return func(ctx context.Context, session DAO) error {
		return err
	}
}

// MockCreateToken mocks a CreateToken function
func MockCreateToken(token string, expiresAt time.Time, err error) CreateToken {
	return func(ctx context.Context, userID int) (string, time.Time, error) {
		return token, expiresAt, err
	}
}

// MockUserSessionDAO mocks a session DAO
func MockUserSessionDAO() DAO {
	return DAO{
		UserID:    1,
		Token:     "abcd1234",
		ExpiresAt: time.Date(2006, time.January, 1, 0, 0, 0, 0, time.Local),
	}
}
