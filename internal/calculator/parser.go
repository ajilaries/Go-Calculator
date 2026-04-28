package calculator

import (
	"strconv"
	"strings"

	"go-calculator/internal/models"
)

func Parse(input string) (models.Expression, error) {
	parts := strings.Fields(input)

	// unary (sqrt 25)
	if len(parts) == 2 {
		val, err := strconv.ParseFloat(parts[1], 64)
		if err != nil {
			return models.Expression{}, ErrInvalidOp
		}

		return models.Expression{
			Operator: parts[0],
			Operands: []float64{val},
		}, nil
	}

	// binary (10 + 5)
	if len(parts) == 3 {
		a, err1 := strconv.ParseFloat(parts[0], 64)
		b, err2 := strconv.ParseFloat(parts[2], 64)

		if err1 != nil || err2 != nil {
			return models.Expression{}, ErrInvalidOp
		}

		return models.Expression{
			Operator: parts[1],
			Operands: []float64{a, b},
		}, nil
	}

	return models.Expression{}, ErrInvalidOp
}