package session_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"

	"github.com/lhbelfanti/corpus-creator/cmd/api/user/session"
)

func TestMain(m *testing.M) {
	_ = godotenv.Load()
	os.Exit(m.Run())
}

func TestCreateToken_success(t *testing.T) {
	mockInsertUserSession := session.MockInsertUserSession(nil)

	createSessionToken := session.MakeCreateToken(mockInsertUserSession)

	token, createdAt, got := createSessionToken(context.Background(), 1)

	assert.Nil(t, got)
	assert.NotNil(t, createdAt)
	assert.NotNil(t, token)
}

func TestCreateToken_failsWhenInsertUserSessionThrowsError(t *testing.T) {
	mockInsertUserSession := session.MockInsertUserSession(errors.New("failed to insert user session"))

	createSessionToken := session.MakeCreateToken(mockInsertUserSession)

	want := session.FailedToCreatUserSessionToken
	_, _, got := createSessionToken(context.Background(), 1)

	assert.Equal(t, want, got)
}
