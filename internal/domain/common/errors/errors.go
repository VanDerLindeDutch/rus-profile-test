package common

import (
	"google.golang.org/grpc/codes"
	"rus-profile-test/internal/domain/common/errors/error_type"
)

type Error struct {
	Code       error_type.ErrorType
	GrpcCode   codes.Code
	InnerError error
}

func (e Error) Error() string {
	return e.InnerError.Error()
}

func NewError(code error_type.ErrorType, grpcCode codes.Code, e error) *Error {

	return &Error{
		Code:       code,
		GrpcCode:   grpcCode,
		InnerError: e,
	}
}
