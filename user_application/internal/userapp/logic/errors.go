package logic

import "errors"

var (
	ErrUserNoFound  = errors.New("user not found")
	ErrInvalidInput = errors.New("input is not valid")
)

type endUserMsgError struct {
	error
	endUserMessage string
}

func NewEndUserMsgError(msg string) error {
	return &endUserMsgError{
		error:          errors.New(msg),
		endUserMessage: msg,
	}
}

func SetEndUserMsg(err error, msg string) error {
	if err == nil {
		return NewEndUserMsgError(msg)
	}

	return &endUserMsgError{
		error:          err,
		endUserMessage: msg,
	}
}

func (e *endUserMsgError) Unwrap() error {
	return e.error
}

func GetEndUserMsg(err error) (string, bool) {
	if endUserMsgErr, ok := errors.AsType[*endUserMsgError](err); ok {
		return endUserMsgErr.endUserMessage, true
	}

	return "", false
}
