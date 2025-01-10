package users

import "errors"

var (
	FailedToInsertUser                  = errors.New("failed to insert user")
	FailedToRetrieveIfUserAlreadyExists = errors.New("failed to retrieve if user already exists")
)
