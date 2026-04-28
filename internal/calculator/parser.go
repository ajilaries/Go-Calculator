package calculator

import (
	"strconv"
	"strings"

	"go-calculator/internal/models"
)

func Parse(input string) (models.Expression, error) {
	parts := strings.Fields(input)

	if len(parts) != 3 {
		return models.Expression{}, ErrInvalidOp
	}

	a, err1 := strconv.ParseFloat(parts[0], 64)
	b, err2 := strconv.ParseFloat(parts[2], 64)

	if err1 != nil || err2 != nil {
		return models.Expression{}, ErrInvalidOp
	}

	return models.Expression{
		A:  a,
		B:  b,
		Op: parts[1],
	}, nil
}