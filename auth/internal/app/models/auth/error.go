package auth

import "errors"

var (
	//ErrAuthAlreadyExists Пользователь уже существует
	ErrAuthAlreadyExists = errors.New("user already exists")
	//ErrAuthInvalidArgument Некорректные аргументы
	ErrAuthInvalidArgument = errors.New("invalid argument")
	//ErrAuthUnauthenticated Не прошел аутентификацию
	ErrAuthUnauthenticated = errors.New("unauthenticated user")
)
