package calculator

import (
	"errors"
	"math"
)

func init() {
	Register("+", add)
	Register("-", subtract)
	Register("*", multiply)
	Register("/", divide)
	Register("^", power)
	Register("%", mod)

	// new
	Register("sqrt", sqrt)
	Register("log", log)
}

func add(nums ...float64) (float64, error) {
	return nums[0] + nums[1], nil
}

func subtract(nums ...float64) (float64, error) {
	return nums[0] - nums[1], nil
}

func multiply(nums ...float64) (float64, error) {
	return nums[0] * nums[1], nil
}

func divide(nums ...float64) (float64, error) {
	if nums[1] == 0 {
		return 0, ErrDivideByZero
	}
	return nums[0] / nums[1], nil
}

func power(nums ...float64) (float64, error) {
	return math.Pow(nums[0], nums[1]), nil
}

func mod(nums ...float64) (float64, error) {
	return math.Mod(nums[0], nums[1]), nil
}
func sqrt(nums ...float64) (float64, error) {
	if nums[0] < 0 {
		return 0, errors.New("invalid sqrt")
	}
	return math.Sqrt(nums[0]), nil
}

func log(nums ...float64) (float64, error) {
	if nums[0] <= 0 {
		return 0, errors.New("invalid log")
	}
	return math.Log(nums[0]), nil
}