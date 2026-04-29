package calculator

import (
	"regexp"
	"strconv"

	"go-calculator/internal/models"
)

func Parse(input string) (models.Expression, error) {
	// Match: number operator number
	re := regexp.MustCompile(`^\s*([\d\.]+)\s*([\+\-\*/\^%])\s*([\d\.]+)\s*$`)
	match := re.FindStringSubmatch(input)

	// binary operation
	if len(match) == 4 {
		a, err1 := strconv.ParseFloat(match[1], 64)
		b, err2 := strconv.ParseFloat(match[3], 64)

		if err1 != nil || err2 != nil {
			return models.Expression{}, ErrInvalidOp
		}

		return models.Expression{
			Operator: match[2],
			Operands: []float64{a, b},
		}, nil
	}

	// unary operation (sqrt 25)
	re2 := regexp.MustCompile(`^\s*(sqrt|log)\s*([\d\.]+)\s*$`)
	match2 := re2.FindStringSubmatch(input)

	if len(match2) == 3 {
		val, err := strconv.ParseFloat(match2[2], 64)
		if err != nil {
			return models.Expression{}, ErrInvalidOp
		}

		return models.Expression{
			Operator: match2[1],
			Operands: []float64{val},
		}, nil
	}

	return models.Expression{}, ErrInvalidOp
}