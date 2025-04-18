package auth

import "errors"

var (
	FailedToRetrieveIfTheUserExists           = errors.New("failed to retrieve user exists")
	FailedToSignUpBecauseTheUserAlreadyExists = errors.New("user already exists")
	FailedToGenerateHashFromPassword          = errors.New("failed to generate hash from password")
	FailedToInsertUserIntoDatabase            = errors.New("failed to insert user into database")
	MissingUsername                           = errors.New("missing username")
	MissingPassword                           = errors.New("missing password")
	FailedToSelectUserByUsername              = errors.New("failed to execute select user by username")
	FailedToLoginDueWrongPassword             = errors.New("failed to login due wrong password")
	FailedToCreateUserSession                 = errors.New("failed to create user session")
	FailedToDeleteUserSession                 = errors.New("failed to delete user session")
	AuthorizationTokenIsRequired              = errors.New("authorization token is required")
)

const (
	InvalidRequestBody         string = "Invalid request body"
	FailedToSignUp             string = "Failed to sign up"
	FailedToLogIn              string = "Failed to log in"
	FailedToLogOut             string = "Failed to log out"
	AuthorizationTokenRequired string = "Authorization token is required"
)
