package auth_test

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

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

func TestLoginHandlerV1_success(t *testing.T) {
	mockExpiresAt := time.Date(2006, time.January, 1, 0, 0, 0, 0, time.Local)
	mockLogIn := auth.MockLogIn("abcd", mockExpiresAt, nil)
	mockResponseWriter := httptest.NewRecorder()
	mockUser := user.MockDTO()
	mockBody, _ := json.Marshal(mockUser)
	mockRequest, _ := http.NewRequestWithContext(context.Background(), http.MethodPost, "/auth/login/v1", bytes.NewReader(mockBody))

	logInHandlerV1 := auth.LogInHandlerV1(mockLogIn)

	logInHandlerV1(mockResponseWriter, mockRequest)

	want := http.StatusOK
	got := mockResponseWriter.Result().StatusCode

	assert.Equal(t, want, got)
}

func TestLoginHandlerV1_failsWhenTheBodyCantBeParsed(t *testing.T) {
	mockExpiresAt := time.Date(2006, time.January, 1, 0, 0, 0, 0, time.Local)
	mockLogIn := auth.MockLogIn("abcd", mockExpiresAt, nil)
	mockResponseWriter := httptest.NewRecorder()
	mockBody, _ := json.Marshal(`{"wrong": "body"}`)
	mockRequest, _ := http.NewRequestWithContext(context.Background(), http.MethodPost, "/auth/login/v1", bytes.NewReader(mockBody))

	logInHandlerV1 := auth.LogInHandlerV1(mockLogIn)

	logInHandlerV1(mockResponseWriter, mockRequest)

	want := http.StatusBadRequest
	got := mockResponseWriter.Result().StatusCode

	assert.Equal(t, want, got)
}

func TestLoginHandlerV1_failsWhenValidateBodyThrowsError(t *testing.T) {
	mockExpiresAt := time.Date(2006, time.January, 1, 0, 0, 0, 0, time.Local)
	mockLogIn := auth.MockLogIn("abcd", mockExpiresAt, nil)
	mockResponseWriter := httptest.NewRecorder()

	for _, test := range []struct {
		mockUser user.DTO
	}{
		{mockUser: user.DTO{Username: "username"}},
		{mockUser: user.DTO{Password: "password"}},
	} {
		mockBody, _ := json.Marshal(test.mockUser)
		mockRequest, _ := http.NewRequestWithContext(context.Background(), http.MethodPost, "/auth/login/v1", bytes.NewReader(mockBody))

		logInHandlerV1 := auth.LogInHandlerV1(mockLogIn)

		logInHandlerV1(mockResponseWriter, mockRequest)

		want := http.StatusBadRequest
		got := mockResponseWriter.Result().StatusCode

		assert.Equal(t, want, got)
	}
}

func TestLoginHandlerV1_failsWhenLogInThrowsError(t *testing.T) {
	for _, test := range []struct {
		logInError error
		want       int
	}{
		{logInError: auth.FailedToLoginDueWrongPassword, want: http.StatusUnauthorized},
		{logInError: errors.New("failed to log in"), want: http.StatusInternalServerError},
	} {
		mockLogIn := auth.MockLogIn("", time.Time{}, test.logInError)
		mockResponseWriter := httptest.NewRecorder()
		mockUser := user.MockDTO()
		mockBody, _ := json.Marshal(mockUser)
		mockRequest, _ := http.NewRequestWithContext(context.Background(), http.MethodPost, "/auth/login/v1", bytes.NewReader(mockBody))

		logInHandlerV1 := auth.LogInHandlerV1(mockLogIn)

		logInHandlerV1(mockResponseWriter, mockRequest)

		want := test.want
		got := mockResponseWriter.Result().StatusCode

		assert.Equal(t, want, got)
	}

}

func TestLogoutHandlerV1_success(t *testing.T) {
	mockLogOut := auth.MockLogout(nil)
	mockResponseWriter := httptest.NewRecorder()
	mockRequest, _ := http.NewRequestWithContext(context.Background(), http.MethodPost, "/auth/logout/v1", nil)
	mockRequest.Header.Set("X-Session-Token", "token")

	logOutHandlerV1 := auth.LogOutHandlerV1(mockLogOut)

	logOutHandlerV1(mockResponseWriter, mockRequest)

	want := http.StatusOK
	got := mockResponseWriter.Result().StatusCode

	assert.Equal(t, want, got)
}

func TestLogoutHandlerV1_failsWhenSessionTokenHeaderWasNotFound(t *testing.T) {
	mockLogOut := auth.MockLogout(nil)
	mockResponseWriter := httptest.NewRecorder()
	mockRequest, _ := http.NewRequestWithContext(context.Background(), http.MethodPost, "/auth/logout/v1", nil)

	logOutHandlerV1 := auth.LogOutHandlerV1(mockLogOut)

	logOutHandlerV1(mockResponseWriter, mockRequest)

	want := http.StatusUnauthorized
	got := mockResponseWriter.Result().StatusCode

	assert.Equal(t, want, got)
}

func TestLogoutHandlerV1_failsWhenLogOutThrowsError(t *testing.T) {
	mockLogOut := auth.MockLogout(errors.New("failed to logout"))
	mockResponseWriter := httptest.NewRecorder()
	mockRequest, _ := http.NewRequestWithContext(context.Background(), http.MethodPost, "/auth/logout/v1", nil)
	mockRequest.Header.Set("X-Session-Token", "token")

	logOutHandlerV1 := auth.LogOutHandlerV1(mockLogOut)

	logOutHandlerV1(mockResponseWriter, mockRequest)

	want := http.StatusInternalServerError
	got := mockResponseWriter.Result().StatusCode

	assert.Equal(t, want, got)
}
