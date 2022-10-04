package uaerr

import "errors"

type userMsgError struct {
	userMsg    string
	wrappedErr error
}

// Boilerplate for implementing "error"
func (u *userMsgError) Error() string {
	if u.wrappedErr == nil {
		return ""
	}

	return u.wrappedErr.Error()
}

// Boilerplate for working with errors.As & errors.Is
func (u *userMsgError) Unwrap() error {
	return u.wrappedErr
}

func NewUserMsg(userMsg string) error {
	return &userMsgError{
		userMsg:    userMsg,
		wrappedErr: errors.New(userMsg),
	}
}

func SetUserMsg(err error, userMsg string) error {
	return &userMsgError{
		userMsg:    userMsg,
		wrappedErr: err,
	}
}

func GetUserMsg(err error) (string, bool) {
	var userMsgErr *userMsgError
	if errors.As(err, &userMsgErr) {
		return userMsgErr.userMsg, true
	}

	return "", false
}
