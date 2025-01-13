package auth_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"ahbcc/cmd/api/auth"
	"ahbcc/cmd/api/user"
	"ahbcc/cmd/api/user/session"
)

func TestLogIn_success(t *testing.T) {
	mockUserDAO := user.MockDAO()
	mockSelectUserByUsername := user.MockSelectByUsername(mockUserDAO, nil)
	mockToken := "abcd"
	mockExpiresAt := time.Date(2006, time.January, 1, 0, 0, 0, 0, time.Local)
	mockCreateSessionToken := session.MockCreateToken(mockToken, mockExpiresAt, nil)
	mockUserDTO := user.MockDTO()

	logIn := auth.MakeLogIn(mockSelectUserByUsername, mockCreateSessionToken)

	token, expiresAt, err := logIn(context.Background(), mockUserDTO)

	assert.Nil(t, err)
	assert.Equal(t, mockToken, token)
	assert.Equal(t, mockExpiresAt, expiresAt)
}

func TestLogIn_failsWhenSelectUserByUsernameThrowsError(t *testing.T) {
	mockUserDAO := user.MockDAO()
	mockSelectUserByUsername := user.MockSelectByUsername(mockUserDAO, errors.New("error while executing SelectByUsername"))
	mockToken := "abcd"
	mockCreatedAt := time.Date(2006, time.January, 1, 0, 0, 0, 0, time.Local)
	mockCreateSessionToken := session.MockCreateToken(mockToken, mockCreatedAt, nil)
	mockUserDTO := user.MockDTO()

	logIn := auth.MakeLogIn(mockSelectUserByUsername, mockCreateSessionToken)

	want := auth.FailedToSelectUserByUsername
	_, _, got := logIn(context.Background(), mockUserDTO)

	assert.Equal(t, want, got)
}

func TestLogIn_failsWhenCompareHashAndPasswordThrowsError(t *testing.T) {
	mockUserDAO := user.MockDAO()
	mockSelectUserByUsername := user.MockSelectByUsername(mockUserDAO, nil)
	mockCreatedAt := time.Date(2006, time.January, 1, 0, 0, 0, 0, time.Local)
	mockCreateSessionToken := session.MockCreateToken("abcd", mockCreatedAt, nil)
	mockUserDTO := user.MockDTO()
	mockUserDTO.Password = "wrong password"

	logIn := auth.MakeLogIn(mockSelectUserByUsername, mockCreateSessionToken)

	want := auth.FailedToLoginDueWrongPassword
	_, _, got := logIn(context.Background(), mockUserDTO)

	assert.Equal(t, want, got)
}

func TestLogIn_failsWhenCreateSessionTokenThrowsError(t *testing.T) {
	mockUserDAO := user.MockDAO()
	mockSelectUserByUsername := user.MockSelectByUsername(mockUserDAO, nil)
	mockCreateSessionToken := session.MockCreateToken("abcd", time.Time{}, errors.New("error while executing CreateSessionToken"))
	mockUserDTO := user.MockDTO()

	logIn := auth.MakeLogIn(mockSelectUserByUsername, mockCreateSessionToken)

	want := auth.FailedToCreateUserSession
	_, _, got := logIn(context.Background(), mockUserDTO)

	assert.Equal(t, want, got)
}
