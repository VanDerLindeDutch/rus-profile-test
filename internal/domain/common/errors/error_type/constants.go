package error_type

import "google.golang.org/grpc/codes"

type ErrorType string

const (
	IncorrectInn ErrorType = "incorrect_inn"
	NotFound     ErrorType = "company_not_found"
	Internal     ErrorType = "internal"
)

func (e ErrorType) GetGrpcCode() codes.Code {
	switch e {
	case IncorrectInn:
		return codes.InvalidArgument
	case NotFound:
		return codes.NotFound
	case Internal:
		return codes.Internal
	}
	return codes.Internal
}
