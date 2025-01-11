package auth_test

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"

	"ahbcc/cmd/api/auth"
	"ahbcc/cmd/api/users"
)

func TestSignUp_success(t *testing.T) {
	mockUserExists := users.MockUserExists(false, nil)
	mockInsertUser := users.MockInsert(nil)
	mockUserDTO := users.MockUserDTO()

	signUp := auth.MakeSignUp(mockUserExists, mockInsertUser)

	got := signUp(context.Background(), mockUserDTO)

	assert.Nil(t, got)
}

func TestSignUp_failsWhenUserExistsThrowsError(t *testing.T) {
	mockUserExists := users.MockUserExists(false, errors.New("failed to execute UserExists"))
	mockInsertUser := users.MockInsert(nil)
	mockUserDTO := users.MockUserDTO()

	signUp := auth.MakeSignUp(mockUserExists, mockInsertUser)

	want := auth.FailedToRetrieveIfTheUserExists
	got := signUp(context.Background(), mockUserDTO)

	assert.Equal(t, want, got)
}

func TestSignUp_failsWhenUserAlreadyExists(t *testing.T) {
	mockUserExists := users.MockUserExists(true, nil)
	mockInsertUser := users.MockInsert(nil)
	mockUserDTO := users.MockUserDTO()

	signUp := auth.MakeSignUp(mockUserExists, mockInsertUser)

	want := auth.FailedToSignUpBecauseTheUserAlreadyExists
	got := signUp(context.Background(), mockUserDTO)

	assert.Equal(t, want, got)
}

func TestSignUp_failsWhenGenerateFromPasswordThrowsError(t *testing.T) {
	mockUserExists := users.MockUserExists(false, nil)
	mockInsertUser := users.MockInsert(nil)
	mockUserDTO := users.MockUserDTO()
	mockUserDTO.Password = "verylongpassword1234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890"

	signUp := auth.MakeSignUp(mockUserExists, mockInsertUser)

	want := auth.FailedToGenerateHashFromPassword
	got := signUp(context.Background(), mockUserDTO)

	assert.Equal(t, want, got)
}

func TestSignUp_failsWhenInsertThrowsError(t *testing.T) {
	mockUserExists := users.MockUserExists(false, nil)
	mockInsertUser := users.MockInsert(errors.New("failed to insert user"))
	mockUserDTO := users.MockUserDTO()

	signUp := auth.MakeSignUp(mockUserExists, mockInsertUser)

	want := auth.FailedToInsertUserIntoDatabase
	got := signUp(context.Background(), mockUserDTO)

	assert.Equal(t, want, got)
}
