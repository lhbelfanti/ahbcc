package session

import "errors"

var (
	FailedToInsertUserSession            = errors.New("failed to insert user session")
	FailedToCreatUserSessionToken        = errors.New("failed to create user session token")
	FailedToDeleteUserSession            = errors.New("failed to delete user session")
	FailedToDeleteExpiredSessions        = errors.New("failed to delete expired sessions")
	NoUserIDFoundForTheGivenToken        = errors.New("no user id found for the given token")
	FailedToExecuteQueryToRetrieveUserID = errors.New("failed to execute query to retrieve user id")
)
