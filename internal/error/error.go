package myerror

import "github.com/cockroachdb/errors"

var (
	ErrVariableNotFound   = errors.New("variable not found")
	ErrUnknownOperation   = errors.New("unknown operation")
	ErrUnknownVariable    = errors.New("unknown variable")
	ErrSomethingWentWrong = errors.New("something went wrong. Check logs")
	ErrDivideByZero       = errors.New("divide by zero")
)
