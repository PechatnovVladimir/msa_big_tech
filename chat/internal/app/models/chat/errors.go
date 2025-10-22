package chat

import "errors"

var (
	ErrNotFound         = errors.New("chat not found")
	ErrAlreadyExists    = errors.New("chat already exists")
	ErrInvalidArgument  = errors.New("chat argument invalid")
	ErrPermissionDenied = errors.New("chat permission denied")
	ErrUnauthenticated  = errors.New("chat unauthenticated")
)
