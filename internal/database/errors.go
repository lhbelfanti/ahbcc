package database

import "errors"

var (
	FailedToPrepareStatement = errors.New("failed to prepare statement")
)
