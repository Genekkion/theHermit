package modal

import "errors"

var (
	// ErrMissingParent is returned when the parent is nil.
	ErrMissingParent = errors.New("missing parent")
	// ErrMissingChild is returned when the child is nil.
	ErrMissingChild = errors.New("missing child")
)
