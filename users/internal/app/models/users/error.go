package users

import "errors"

var (
	// ErrUserNotFound Пользователь не найден
	ErrUserNotFound = errors.New("user not found")
	//ErrUserAlreadyExists Пользователь уже существует
	ErrUserAlreadyExists = errors.New("user already exists")
	//ErrUserInvalidArgument Некорректные аргументы
	ErrUserInvalidArgument = errors.New("invalid argument")
)
