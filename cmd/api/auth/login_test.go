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

func TestLogin_success(t *testing.T) {
	mockUserDAO := user.MockDAO()
	mockSelectUserByUsername := user.MockSelectByUsername(mockUserDAO, nil)
	mockToken := "abcd"
	mockCreatedAt := time.Date(2006, time.January, 1, 0, 0, 0, 0, time.Local)
	mockCreateSessionToken := session.MockCreateToken(mockToken, mockCreatedAt, nil)
	mockUserDTO := user.MockDTO()

	login := auth.MakeLogin(mockSelectUserByUsername, mockCreateSessionToken)

	token, createdAt, err := login(context.Background(), mockUserDTO)

	assert.Nil(t, err)
	assert.Equal(t, mockToken, token)
	assert.Equal(t, mockCreatedAt, createdAt)
}

func TestLogin_failsWhenSelectUserByUsernameThrowsError(t *testing.T) {
	mockUserDAO := user.MockDAO()
	mockSelectUserByUsername := user.MockSelectByUsername(mockUserDAO, errors.New("error while executing SelectByUsername"))
	mockToken := "abcd"
	mockCreatedAt := time.Date(2006, time.January, 1, 0, 0, 0, 0, time.Local)
	mockCreateSessionToken := session.MockCreateToken(mockToken, mockCreatedAt, nil)
	mockUserDTO := user.MockDTO()

	login := auth.MakeLogin(mockSelectUserByUsername, mockCreateSessionToken)

	want := auth.FailedToSelectUserByUsername
	_, _, got := login(context.Background(), mockUserDTO)

	assert.Equal(t, want, got)
}

func TestLogin_failsWhenCompareHashAndPasswordThrowsError(t *testing.T) {
	mockUserDAO := user.MockDAO()
	mockSelectUserByUsername := user.MockSelectByUsername(mockUserDAO, nil)
	mockCreatedAt := time.Date(2006, time.January, 1, 0, 0, 0, 0, time.Local)
	mockCreateSessionToken := session.MockCreateToken("abcd", mockCreatedAt, nil)
	mockUserDTO := user.MockDTO()
	mockUserDTO.Password = "wrong password"

	login := auth.MakeLogin(mockSelectUserByUsername, mockCreateSessionToken)

	want := auth.FailedToLoginDueWrongPassword
	_, _, got := login(context.Background(), mockUserDTO)

	assert.Equal(t, want, got)
}

func TestLogin_failsWhenCreateSessionTokenThrowsError(t *testing.T) {
	mockUserDAO := user.MockDAO()
	mockSelectUserByUsername := user.MockSelectByUsername(mockUserDAO, nil)
	mockCreateSessionToken := session.MockCreateToken("abcd", time.Time{}, errors.New("error while executing CreateSessionToken"))
	mockUserDTO := user.MockDTO()

	login := auth.MakeLogin(mockSelectUserByUsername, mockCreateSessionToken)

	want := auth.FailedToCreateUserSession
	_, _, got := login(context.Background(), mockUserDTO)

	assert.Equal(t, want, got)
}
