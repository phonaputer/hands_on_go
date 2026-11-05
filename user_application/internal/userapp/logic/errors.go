package logic

import (
	"errors"
)

// 1. Error variables
// Pros
//   * Error comparison - errors.Is(err, ErrUserNotFound)
//   * These are standard for Go, everyone understands this pattern easily
//   * easy to understand pattern, it's just a variable
//
// Cons
//   * It's just a variable, it can't contain additional information beyond a message
//
//

var (
	ErrUserNotFound = errors.New("user not found")
	ErrInvalidInput = errors.New("input is not valid")
)

// 2. Custom `error` implementations
// Pros
//   * Now you can have additional information beyond the message
//
// Cons
//   * It's a bit more boilerplate
//   * It's a more complex pattern
//   * Unless you're also using error variables, standard library `errors.Is` no longer will work for you

type endUserMsgError struct {
	error
	endUserMessage string
}

func (e *endUserMsgError) Unwrap() error {
	return e.error
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
