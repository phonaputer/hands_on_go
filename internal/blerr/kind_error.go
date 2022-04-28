package blerr

import "errors"

// kindError is an error which contains a business logic "kind"
type kindError struct {
	kind Kind
	err  error
}

func (e *kindError) Error() string {
	if e.err == nil {
		return ""
	}

	return e.err.Error()
}

func (e *kindError) Unwrap() error {
	return e.err
}

func SetKind(err error, kind Kind) error {
	return &kindError{
		kind: kind,
		err:  err,
	}
}

func GetKind(err error) Kind {
	var kindErr *kindError
	if errors.As(err, &kindErr) {
		return kindErr.kind
	}

	return KindUnknown
}
