package models

import "errors"

var (
	ErrSocialInvalidArgument  = errors.New("invalid argument")
	ErrSocialNotFound         = errors.New("not found")
	ErrSocialAlreadyExists    = errors.New("already exists")
	ErrSocialPermissionDenied = errors.New("permission denied")
)
