package user

import "errors"

var (
	FailedToInsertUser                  = errors.New("failed to insert user")
	FailedToRetrieveIfUserAlreadyExists = errors.New("failed to retrieve if user already exists")
	NoUserFoundForTheGivenUsername      = errors.New("no user found for the given username")
	FailedExecuteQueryToRetrieveUser    = errors.New("failed to execute query to retrieve user")
)
