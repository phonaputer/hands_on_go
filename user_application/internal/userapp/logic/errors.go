package logic

import (
	"errors"
)

var (
	ErrUserNotFound = errors.New("user not found")
	ErrInvalidInput = errors.New("input is not valid")
)

type endUserMsgError struct {
	error
	endUserMessage string
}

func NewEndUserMsgErr(msg string) error {
	return &endUserMsgError{
		error:          errors.New(msg),
		endUserMessage: msg,
	}
}

func SetEndUserMsg(err error, msg string) error {
	if err == nil {
		return NewEndUserMsgErr(msg)
	}

	return &endUserMsgError{
		error:          err,
		endUserMessage: msg,
	}
}

func GetEndUserMsg(err error) (string, bool) {
	if endUserMsgErr, ok := errors.AsType[*endUserMsgError](err); ok {
		return endUserMsgErr.endUserMessage, true
	}

	return "", false
}
