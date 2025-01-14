package auth

import (
	"encoding/json"
	"errors"
	"net/http"

	"ahbcc/cmd/api/user"
	"ahbcc/internal/log"
)

// SignUpHandlerV1 HTTP Handler of the endpoint /auth/signup/v1
func SignUpHandlerV1(signUp SignUp) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		var userDTO user.DTO
		err := json.NewDecoder(r.Body).Decode(&userDTO)
		if err != nil {
			log.Error(ctx, err.Error())
			http.Error(w, InvalidRequestBody, http.StatusBadRequest)
			return
		}
		ctx = log.With(ctx, log.Param("username", userDTO.Username))

		err = validateBody(userDTO)
		if err != nil {
			log.Error(ctx, err.Error())
			http.Error(w, InvalidRequestBody, http.StatusBadRequest)
		}

		err = signUp(ctx, userDTO)
		if err != nil {
			log.Error(ctx, err.Error())
			http.Error(w, FailedToSignUp, http.StatusInternalServerError)
			return
		}

		log.Info(ctx, "User successfully signed up")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("User successfully signed up"))
	}
}

// LogInHandlerV1 HTTP Handler of the endpoint /auth/login/v1
func LogInHandlerV1(logIn LogIn) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		var userDTO user.DTO
		err := json.NewDecoder(r.Body).Decode(&userDTO)
		if err != nil {
			log.Error(ctx, err.Error())
			http.Error(w, InvalidRequestBody, http.StatusBadRequest)
			return
		}

		err = validateBody(userDTO)
		if err != nil {
			log.Error(ctx, err.Error())
			http.Error(w, InvalidRequestBody, http.StatusBadRequest)
		}

		token, expiresAt, err := logIn(ctx, userDTO)
		if err != nil {
			log.Error(ctx, err.Error())

			switch {
			case errors.Is(err, FailedToLoginDueWrongPassword):
				http.Error(w, FailedToLogIn, http.StatusUnauthorized)
				return
			default:
				http.Error(w, FailedToLogIn, http.StatusInternalServerError)
				return
			}
		}

		loginResponse := LoginResponse{
			Token:     token,
			ExpiresAt: expiresAt,
		}

		log.Info(ctx, "User successfully logged in")
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(loginResponse)
	}
}

// LogOutHandlerV1 HTTP Handler of the endpoint /auth/logout/v1
func LogOutHandlerV1(logOut LogOut) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		token := r.Header.Get("X-Session-Token")
		if token == "" {
			log.Error(ctx, AuthorizationTokenRequired)
			http.Error(w, AuthorizationTokenRequired, http.StatusUnauthorized)
			return
		}

		err := logOut(ctx, token)
		if err != nil {
			log.Error(ctx, err.Error())
			http.Error(w, FailedToLogOut, http.StatusInternalServerError)
		}

		log.Info(ctx, "User successfully logged out")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("User successfully logged out"))
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
