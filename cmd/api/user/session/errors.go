package session

import "errors"

var (
	FailedToInsertUserSession     = errors.New("failed to insert user session")
	FailedToCreatUserSessionToken = errors.New("failed to create user session token")
	FailedToDeleteUserSession     = errors.New("failed to delete user session")
)
