package users

import "errors"

var (
	//ErrUnauthenticated Пользователь не авторизован
	ErrUnauthenticated = errors.New("user not authenticated")
	//ErrPermissionDenied Доступ запрещен
	ErrPermissionDenied = errors.New("permission denied")
	// ErrUserNotFound Пользователь не найден
	ErrUserNotFound = errors.New("user not found")
	//ErrUserAlreadyExists Пользователь уже существует
	ErrUserAlreadyExists = errors.New("user already exists")
	//ErrUserInvalidArgument Некорректные аргументы
	ErrUserInvalidArgument = errors.New("invalid argument")
)
