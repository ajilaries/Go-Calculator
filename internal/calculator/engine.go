package calculator

// import "go-calculator/internal/models"

type Operation func(nums  ...float64) (float64, error)

var operations = map[string]Operation{}

// Register new operations
func Register(op string, fn Operation) {
	operations[op] = fn
}

// Execute calculation
func Execute(op string, nums ...float64) (float64, error) {
	if fn, ok := operations[op]; ok {
		return fn(nums...)
	}
	return 0, ErrInvalidOp
}