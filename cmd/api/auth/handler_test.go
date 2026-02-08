package auth_test

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/lhbelfanti/corpus-creator/cmd/api/auth"
	"github.com/lhbelfanti/corpus-creator/cmd/api/user"
	"github.com/lhbelfanti/corpus-creator/internal/http/response"
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

func TestSignUpHandlerV1_failsWhenTheBodyCannotBeParsed(t *testing.T) {
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
	mockToken := "abcd"
	mockExpiresAt := time.Date(2006, time.January, 1, 0, 0, 0, 0, time.UTC)
	mockLogIn := auth.MockLogIn(mockToken, mockExpiresAt, nil)
	mockResponseWriter := httptest.NewRecorder()
	mockUser := user.MockDTO()
	mockBody, _ := json.Marshal(mockUser)
	mockRequest, _ := http.NewRequestWithContext(context.Background(), http.MethodPost, "/auth/login/v1", bytes.NewReader(mockBody))

	logInHandlerV1 := auth.LogInHandlerV1(mockLogIn)

	logInHandlerV1(mockResponseWriter, mockRequest)

	body, err := io.ReadAll(mockResponseWriter.Result().Body)
	if err != nil {
		t.Fatalf("Failed to read response body: %v", err)
	}

	want := auth.LoginResponseDTO{Token: mockToken, ExpiresAt: mockExpiresAt}

	var responseDTO response.DTO
	err = json.Unmarshal(body, &responseDTO)
	if err != nil {
		t.Fatalf("Failed to parse response body as JSON: %v", err)
	}

	var got auth.LoginResponseDTO
	dataBytes, err := json.Marshal(responseDTO.Data)
	if err != nil {
		t.Fatalf("Failed to marshal Data field: %v", err)
	}

	err = json.Unmarshal(dataBytes, &got)
	if err != nil {
		t.Fatalf("Failed to unmarshal Data field into LoginResponseDTO: %v", err)
	}

	assert.Equal(t, want, got)
	assert.Equal(t, "application/json", mockResponseWriter.Header().Get("Content-Type"))
	assert.Equal(t, http.StatusOK, mockResponseWriter.Result().StatusCode)
}

func TestLoginHandlerV1_failsWhenTheBodyCannotBeParsed(t *testing.T) {
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
