package dto

import "errors"

var (
	ErrChatNotFound         = errors.New("chat not found")
	ErrChatAlreadyExists    = errors.New("chat already exists")
	ErrChatInvalidArgument  = errors.New("chat argument invalid")
	ErrChatPermissionDenied = errors.New("chat permission denied")
)
