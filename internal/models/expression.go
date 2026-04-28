package models

type Expression struct {
	Operator string
	Operands []float64
	RawInput string   // original input
}