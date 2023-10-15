package errs

import "errors"

var (
	// ErrInvalidDatabaseOperation
	// This error is used to describe database operations.
	// These operations may not have any exceptions themselves, but they have not caused any impact, so they are invalid operations.
	ErrInvalidDatabaseOperation = errors.New("invalid database operation")
)
