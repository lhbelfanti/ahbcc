package auth_test

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	"ahbcc/cmd/api/auth"
	"ahbcc/cmd/api/user"
)

func TestSignUpHandlerV1_success(t *testing.T) {
	mockSignUp := auth.MockSignUp(nil)
	mockResponseWriter := httptest.NewRecorder()
	mockUser := user.MockDTO()
	mockBody, _ := json.Marshal(mockUser)
	mockRequest, _ := http.NewRequestWithContext(context.Background(), http.MethodPost, "/auth/signup/v1", bytes.NewReader(mockBody))

	signUpHandlerV1 := auth.SignUpHandlerV1(mockSignUp)

	signUpHandlerV1(mockResponseWriter, mockRequest)

	want := http.StatusOK
	got := mockResponseWriter.Result().StatusCode

	assert.Equal(t, want, got)
}

func TestSignUpHandlerV1_failsWhenTheBodyCantBeParsed(t *testing.T) {
	mockSignUp := auth.MockSignUp(nil)
	mockResponseWriter := httptest.NewRecorder()
	mockBody, _ := json.Marshal(`{"wrong": "body"}`)
	mockRequest, _ := http.NewRequestWithContext(context.Background(), http.MethodPost, "/auth/signup/v1", bytes.NewReader(mockBody))

	signUpHandlerV1 := auth.SignUpHandlerV1(mockSignUp)

	signUpHandlerV1(mockResponseWriter, mockRequest)

	want := http.StatusBadRequest
	got := mockResponseWriter.Result().StatusCode

	assert.Equal(t, want, got)
}

func TestSignUpHandlerV1_failsWhenValidateBodyThrowsError(t *testing.T) {
	mockSignUp := auth.MockSignUp(nil)
	mockResponseWriter := httptest.NewRecorder()

	for _, test := range []struct {
		mockUser user.DTO
	}{
		{mockUser: user.DTO{Username: "username"}},
		{mockUser: user.DTO{Password: "password"}},
	} {
		mockBody, _ := json.Marshal(test.mockUser)
		mockRequest, _ := http.NewRequestWithContext(context.Background(), http.MethodPost, "/auth/signup/v1", bytes.NewReader(mockBody))

		signUpHandlerV1 := auth.SignUpHandlerV1(mockSignUp)

		signUpHandlerV1(mockResponseWriter, mockRequest)

		want := http.StatusBadRequest
		got := mockResponseWriter.Result().StatusCode

		assert.Equal(t, want, got)
	}
}

func TestSignUpHandlerV1_failsWhenSignUpThrowsError(t *testing.T) {
	mockSignUp := auth.MockSignUp(errors.New("failed to sign up"))
	mockResponseWriter := httptest.NewRecorder()
	mockUser := user.MockDTO()
	mockBody, _ := json.Marshal(mockUser)
	mockRequest, _ := http.NewRequestWithContext(context.Background(), http.MethodPost, "/auth/signup/v1", bytes.NewReader(mockBody))

	signUpHandlerV1 := auth.SignUpHandlerV1(mockSignUp)

	signUpHandlerV1(mockResponseWriter, mockRequest)

	want := http.StatusInternalServerError
	got := mockResponseWriter.Result().StatusCode

	assert.Equal(t, want, got)
}
