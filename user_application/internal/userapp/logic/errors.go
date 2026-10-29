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
