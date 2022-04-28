package blerr

import "errors"

// userMsgError is an error which contains a message for display to end users
type userMsgError struct {
	userMsg string
	err     error
}

func (e *userMsgError) Error() string {
	if e.err == nil {
		return ""
	}

	return e.err.Error()
}

func (e *userMsgError) Unwrap() error {
	return e.err
}

func SetUserMsg(err error, userMsg string) error {
	return &userMsgError{
		userMsg: userMsg,
		err:     err,
	}
}

func GetUserMsg(err error) (string, bool) {
	var userMsgErr *userMsgError
	if errors.As(err, &userMsgErr) {
		return userMsgErr.userMsg, true
	}

	return "", false
}
