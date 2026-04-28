package calculator

import "errors"

var(
	ErrDivideByZero=errors.New("cannot divide by Zero")
	ErrInvalidOp=errors.New("Invalid Operations")
)