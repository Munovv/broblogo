package messages

import "errors"

var (
	ErrPostNotFound      = errors.New("posts not found for current user")
	ErrAccessDenied      = errors.New("access denied for current user")
	ErrUserNotFound      = errors.New("user id not found")
	ErrUserIdInvalidType = errors.New("user id is of invalid type")
)
