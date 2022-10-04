package uaerr

import "errors"

type typeError struct {
	errType    Type
	wrappedErr error
}

// Boilerplate for implementing "error"
func (t *typeError) Error() string {
	if t.wrappedErr == nil {
		return ""
	}

	return t.wrappedErr.Error()
}

// Boilerplate for working with errors.As & errors.Is
func (t *typeError) Unwrap() error {
	return t.wrappedErr
}

func SetType(err error, errType Type) error {
	return &typeError{
		errType:    errType,
		wrappedErr: err,
	}
}

func GetType(err error) Type {
	var typeErr *typeError
	if errors.As(err, &typeErr) {
		return typeErr.errType
	}

	return TypeUnknown
}
