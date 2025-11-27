package box

import "errors"

var (
	ErrMissingParent = errors.New("missing parent")
	ErrMissingChild  = errors.New("missing child")
)
