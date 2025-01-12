package session

import "time"

// MockUserSessionDAO mocks a session DAO
func MockUserSessionDAO() DAO {
	return DAO{
		UserID:    1,
		Token:     "abcd1234",
		ExpiresAt: time.Date(2006, time.January, 1, 0, 0, 0, 0, time.Local),
	}
}
