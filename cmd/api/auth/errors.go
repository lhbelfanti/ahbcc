package auth

import "errors"

var (
	FailedToRetrieveIfTheUserExists           = errors.New("failed to retrieve user exists")
	FailedToSignUpBecauseTheUserAlreadyExists = errors.New("user already exists")
	FailedToGenerateHashFromPassword          = errors.New("failed to generate hash from password")
	FailedToInsertUserIntoDatabase            = errors.New("failed to insert user into database")
	MissingUsername                           = errors.New("missing username")
	MissingPassword                           = errors.New("missing password")
)

const (
	InvalidRequestBody = "Invalid request body"
	FailedToSignUp     = "Failed to sign up"
)
