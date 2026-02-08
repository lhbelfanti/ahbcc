package auth

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/lhbelfanti/corpus-creator/cmd/api/user"
	"github.com/lhbelfanti/corpus-creator/internal/http/response"
	"github.com/lhbelfanti/corpus-creator/internal/log"
)

// SignUpHandlerV1 HTTP Handler of the endpoint /auth/signup/v1
func SignUpHandlerV1(signUp SignUp) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		var userDTO user.DTO
		err := json.NewDecoder(r.Body).Decode(&userDTO)
		if err != nil {
			response.Send(ctx, w, http.StatusBadRequest, InvalidRequestBody, nil, err)
			return
		}
		ctx = log.With(ctx, log.Param("username", userDTO.Username))

		err = validateBody(userDTO)
		if err != nil {
			response.Send(ctx, w, http.StatusBadRequest, InvalidRequestBody, nil, err)
			return
		}

		err = signUp(ctx, userDTO)
		if err != nil {
			response.Send(ctx, w, http.StatusInternalServerError, FailedToSignUp, nil, err)
			return
		}

		response.Send(ctx, w, http.StatusOK, "User successfully signed up", nil, nil)
	}
}

// LogInHandlerV1 HTTP Handler of the endpoint /auth/login/v1
func LogInHandlerV1(logIn LogIn) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		var userDTO user.DTO
		err := json.NewDecoder(r.Body).Decode(&userDTO)
		if err != nil {
			response.Send(ctx, w, http.StatusBadRequest, InvalidRequestBody, nil, err)
			return
		}

		err = validateBody(userDTO)
		if err != nil {
			response.Send(ctx, w, http.StatusBadRequest, InvalidRequestBody, nil, err)
			return
		}

		token, expiresAt, err := logIn(ctx, userDTO)
		if err != nil {
			switch {
			case errors.Is(err, FailedToLoginDueWrongPassword):
				response.Send(ctx, w, http.StatusUnauthorized, FailedToLogIn, nil, err)
				return
			default:
				response.Send(ctx, w, http.StatusInternalServerError, FailedToLogIn, nil, err)
				return
			}
		}

		loginResponse := LoginResponseDTO{
			Token:     token,
			ExpiresAt: expiresAt,
		}

		response.Send(ctx, w, http.StatusOK, "User successfully logged in", loginResponse, nil)
	}
}

// LogOutHandlerV1 HTTP Handler of the endpoint /auth/logout/v1
func LogOutHandlerV1(logOut LogOut) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		token := r.Header.Get("X-Session-Token")
		if token == "" {
			response.Send(ctx, w, http.StatusUnauthorized, AuthorizationTokenRequired, nil, AuthorizationTokenIsRequired)
			return
		}

		err := logOut(ctx, token)
		if err != nil {
			response.Send(ctx, w, http.StatusInternalServerError, FailedToLogOut, nil, err)
			return
		}

		response.Send(ctx, w, http.StatusOK, "User successfully logged out", nil, nil)
	}
}

// validateBody validates that mandatory fields are present
func validateBody(user user.DTO) error {
	if user.Username == "" {
		return MissingUsername
	}

	if user.Password == "" {
		return MissingPassword
	}

	return nil
}
