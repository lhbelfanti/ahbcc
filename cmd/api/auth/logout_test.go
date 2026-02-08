package auth_test

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/lhbelfanti/corpus-creator/cmd/api/auth"
	"github.com/lhbelfanti/corpus-creator/cmd/api/user/session"
)

func TestLogOut_success(t *testing.T) {
	mockDeleteUserSession := session.MockDelete(nil)

	logOut := auth.MakeLogOut(mockDeleteUserSession)

	got := logOut(context.Background(), "token")

	assert.Nil(t, got)
}

func TestLogOut_failsWhenDeleteSessionThrowsError(t *testing.T) {
	mockDeleteUserSession := session.MockDelete(errors.New("failed to delete session"))

	logOut := auth.MakeLogOut(mockDeleteUserSession)

	want := auth.FailedToDeleteUserSession
	got := logOut(context.Background(), "token")

	assert.Equal(t, want, got)
}
