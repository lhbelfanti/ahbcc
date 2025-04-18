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

// MockDelete mocks a Delete function
func MockDelete(err error) Delete {
	return func(ctx context.Context, token string) error {
		return err
	}
}

// MockDeleteExpiredSessions mocks a DeleteExpiredSessions function
func MockDeleteExpiredSessions(err error) DeleteExpiredSessions {
	return func(ctx context.Context, userID int) error {
		return err
	}
}

// MockSelectUserIDByToken mocks a SelectUserIDByToken function
func MockSelectUserIDByToken(userID int, err error) SelectUserIDByToken {
	return func(ctx context.Context, token string) (int, error) {
		return userID, err
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
