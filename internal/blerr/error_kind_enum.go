package blerr

type Kind int

const (
	KindUnknown Kind = iota
	KindNotFound
	KindInvalidInput
)
