package calculator

import "go-calculator/internal/models"

type Operation func(a, b float64) (float64, error)

var operations = map[string]Operation{}

// Register new operations
func Register(op string, fn Operation) {
	operations[op] = fn
}

// Execute calculation
func Execute(exp models.Expression) (float64, error) {
	if fn, ok := operations[exp.Op]; ok {
		return fn(exp.A, exp.B)
	}
	return 0, ErrInvalidOp
}