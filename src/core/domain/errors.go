package domain

import "errors"

var (
	//System error
	ErrByteReading         = errors.New("unable to read bytes")
	ErrUnmarshallingFailed = errors.New("unable to decode struct")

	ErrBadUserPassword         = errors.New("failed to validate user/password")
	ErrUnableToReachUserServer = errors.New("unable to reach user server")
	ErrUnableToFindUserServer  = errors.New("unable to find user server (found)")
)
