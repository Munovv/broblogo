package errors

import "errors"

var (
	UserNotFound         = errors.New("пользователь не найден")
	InvalidSigningMethod = errors.New("не валидный метод подписи у jwt-токена")
	InvalidTokenType     = errors.New("не валидный тип jwt-токена")
)
