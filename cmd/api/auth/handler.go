package auth

import (
	"ahbcc/internal/log"
	"encoding/json"
	"net/http"

	"ahbcc/cmd/api/users"
)

// SignUpHandlerV1 HTTP Handler of the endpoint /auth/signup
func SignUpHandlerV1(signUp SignUp) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		var user users.UserDTO
		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			log.Error(ctx, err.Error())
			http.Error(w, InvalidRequestBody, http.StatusBadRequest)
			return
		}
		ctx = log.With(ctx, log.Param("username", user.Username))

		err = validateBody(user)
		if err != nil {
			log.Error(ctx, err.Error())
			http.Error(w, InvalidRequestBody, http.StatusBadRequest)
		}

		err = signUp(ctx, user)
		if err != nil {
			log.Error(ctx, err.Error())
			http.Error(w, FailedToSignUp, http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("User successfully signed up"))
	}
}

// validateBody validates that mandatory fields are present
func validateBody(user users.UserDTO) error {
	if user.Username == "" {
		return MissingUsername
	}

	if user.Password == "" {
		return MissingPassword
	}

	return nil
}
