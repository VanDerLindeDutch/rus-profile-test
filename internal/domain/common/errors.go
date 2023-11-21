package common

import (
	"rus-profile-test/internal/domain/common/error_type"
)

type Error struct {
	Code       error_type.ErrorType
	InnerError error
}

func (e Error) Error() string {
	return e.InnerError.Error()
}

func NewError(code error_type.ErrorType, e error) *Error {

	return &Error{
		Code:       code,
		InnerError: e,
	}
}
