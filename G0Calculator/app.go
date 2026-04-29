package main

import (
	"context"
	"fmt"

	"go-calculator/internal/calculator"
)

type App struct {
	ctx context.Context
}

func NewApp() *App {
	return &App{}
}

// 🔥 THIS FIXES YOUR ERROR
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// function accessible from UI
func (a *App) Calculate(input string) (string, error) {
	exp, err := calculator.Parse(input)
	if err != nil {
		return "", err
	}

	res, err := calculator.Execute(exp.Operator, exp.Operands...)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%.4f", res), nil
}