package auth_test

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/lhbelfanti/corpus-creator/cmd/api/auth"
	"github.com/lhbelfanti/corpus-creator/cmd/api/user"
)

func TestSignUp_success(t *testing.T) {
	mockUserExists := user.MockExists(false, nil)
	mockInsertUser := user.MockInsert(nil)
	mockUserDTO := user.MockDTO()

	signUp := auth.MakeSignUp(mockUserExists, mockInsertUser)

	got := signUp(context.Background(), mockUserDTO)

	assert.Nil(t, got)
}

func TestSignUp_failsWhenUserExistsThrowsError(t *testing.T) {
	mockUserExists := user.MockExists(false, errors.New("failed to execute Exists"))
	mockInsertUser := user.MockInsert(nil)
	mockUserDTO := user.MockDTO()

	signUp := auth.MakeSignUp(mockUserExists, mockInsertUser)

	want := auth.FailedToRetrieveIfTheUserExists
	got := signUp(context.Background(), mockUserDTO)

	assert.Equal(t, want, got)
}

func TestSignUp_failsWhenUserAlreadyExists(t *testing.T) {
	mockUserExists := user.MockExists(true, nil)
	mockInsertUser := user.MockInsert(nil)
	mockUserDTO := user.MockDTO()

	signUp := auth.MakeSignUp(mockUserExists, mockInsertUser)

	want := auth.FailedToSignUpBecauseTheUserAlreadyExists
	got := signUp(context.Background(), mockUserDTO)

	assert.Equal(t, want, got)
}

func TestSignUp_failsWhenGenerateFromPasswordThrowsError(t *testing.T) {
	mockUserExists := user.MockExists(false, nil)
	mockInsertUser := user.MockInsert(nil)
	mockUserDTO := user.MockDTO()
	mockUserDTO.Password = "verylongpassword1234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890"

	signUp := auth.MakeSignUp(mockUserExists, mockInsertUser)

	want := auth.FailedToGenerateHashFromPassword
	got := signUp(context.Background(), mockUserDTO)

	assert.Equal(t, want, got)
}

func TestSignUp_failsWhenInsertThrowsError(t *testing.T) {
	mockUserExists := user.MockExists(false, nil)
	mockInsertUser := user.MockInsert(errors.New("failed to insert user"))
	mockUserDTO := user.MockDTO()

	signUp := auth.MakeSignUp(mockUserExists, mockInsertUser)

	want := auth.FailedToInsertUserIntoDatabase
	got := signUp(context.Background(), mockUserDTO)

	assert.Equal(t, want, got)
}
