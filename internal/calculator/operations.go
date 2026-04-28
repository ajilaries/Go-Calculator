package calculator

import "math"

func init() {
	Register("+", add)
	Register("-", subtract)
	Register("*", multiply)
	Register("/", divide)
	Register("^", power)
	Register("%", mod)
}

func add(a, b float64) (float64, error) {
	return a + b, nil
}

func subtract(a, b float64) (float64, error) {
	return a - b, nil
}

func multiply(a, b float64) (float64, error) {
	return a * b, nil
}

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, ErrDivideByZero
	}
	return a / b, nil
}

func power(a, b float64) (float64, error) {
	return math.Pow(a, b), nil
}

func mod(a, b float64) (float64, error) {
	return math.Mod(a, b), nil
}