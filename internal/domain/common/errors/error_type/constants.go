package error_type

type ErrorType string

const (
	IncorrectInn ErrorType = "incorrect_inn"
	NotFound     ErrorType = "not_found"
	Internal     ErrorType = "internal"
)
